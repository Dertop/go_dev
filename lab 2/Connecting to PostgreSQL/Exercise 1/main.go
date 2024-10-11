package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "postgres"
	dbPassword = "123"
	dbName     = "postgres"
)

type Todo struct {
	ID        int
	Task      string
	Completed bool
}

func DBConnect() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to db")
	return db, nil
}

func CreateTodo(db *sql.DB, todo Todo) (int, error) {
	var newID int
	err := db.QueryRow("INSERT INTO todos (task, completed) VALUES ($1, $2) RETURNING id", todo.Task, todo.Completed).Scan(&newID)
	if err != nil {
		return 0, err
	}
	return newID, nil
}

func GetTodo(db *sql.DB, id int) (Todo, error) {
	var todo Todo
	err := db.QueryRow("SELECT id, task, completed FROM todos WHERE id = $1", id).Scan(&todo.ID, &todo.Task, &todo.Completed)
	if err != nil {
		return Todo{}, err
	}
	return todo, nil
}

func UpdateTodo(db *sql.DB, todo Todo) error {
	_, err := db.Exec("UPDATE todos SET task=$1, completed=$2 WHERE id=$3", todo.Task, todo.Completed, todo.ID)
	return err
}

func DeleteTodo(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM todos WHERE id = $1", id)
	return err
}

func GetAllTodos(db *sql.DB) ([]Todo, error) {
	rows, err := db.Query("SELECT id, task, completed FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todoList []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Task, &todo.Completed); err != nil {
			return nil, err
		}
		todoList = append(todoList, todo)
	}

	return todoList, nil
}

func main() {
	db, err := DBConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	task := Todo{Task: "Complete the assignment", Completed: false}
	taskID, err := CreateTodo(db, task)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("New task created with ID: %d\n", taskID)

	retrievedTask, err := GetTodo(db, taskID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Task retrieved: %+v\n", retrievedTask)

	retrievedTask.Task = "Update the task"
	retrievedTask.Completed = true
	if err := UpdateTodo(db, retrievedTask); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Task updated")

	allTasks, err := GetAllTodos(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("All tasks: %+v\n", allTasks)
}
