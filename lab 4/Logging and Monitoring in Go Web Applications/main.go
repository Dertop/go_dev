package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger = logrus.New()

var (
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of request durations",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)
)

func init() {
	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(requestDuration)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &statusResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(wrapper, r)
		duration := time.Since(start).Seconds()

		requestsTotal.WithLabelValues(r.Method, r.URL.Path, http.StatusText(wrapper.statusCode)).Inc()
		requestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)

		logger.WithFields(logrus.Fields{
			"method":   r.Method,
			"path":     r.URL.Path,
			"status":   wrapper.statusCode,
			"duration": duration,
		}).Info("HTTP Request")
	})
}

type statusResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *statusResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Home page accessed successfully")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the logging and monitoring app!"))
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	logger.Warn("Potential issue detected on error page")
	err := simulateError()
	if err != nil {
		logger.WithError(err).Error("Critical error occurred")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("This will never be reached"))
}

func simulateError() error {
	return errors.New("simulated error for demonstration purposes")
}

func main() {
	// Настраиваем логгер для ротации
	logger.SetOutput(&lumberjack.Logger{
		Filename:   "app.log", // Имя файла
		MaxSize:    10,        // Максимальный размер файла в MB
		MaxBackups: 5,         // Сколько файлов резервной копии хранить
		MaxAge:     30,        // Сколько дней хранить логи
		Compress:   true,      // Сжимать резервные файлы
	})
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)

	router := mux.NewRouter()

	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/error", errorHandler).Methods("GET")
	router.Handle("/metrics", promhttp.Handler())
	router.Use(loggingMiddleware)

	logger.Info("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		logger.WithError(err).Fatal("Server failed to start")
	}
}
