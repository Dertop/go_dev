package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "postgres"
	dbPassword = "123"
	dbName     = "postgres"
)

func establishConnection() (*sql.DB, error) {
	connectionInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	database, err := sql.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}

	database.SetMaxOpenConns(15)
	database.SetMaxIdleConns(7)
	database.SetConnMaxLifetime(45 * time.Minute)

	if err = database.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Database connected successfully!")
	return database, nil
}

func initializeTable(db *sql.DB) error {
	createQuery := `
		CREATE TABLE IF NOT EXISTS user_accounts (
			id SERIAL PRIMARY KEY,
			full_name VARCHAR(50) UNIQUE NOT NULL,
			years_old INT NOT NULL
		);`

	_, err := db.Exec(createQuery)
	if err != nil {
		return fmt.Errorf("unable to create table: %w", err)
	}

	fmt.Println("Table 'user_accounts' was created!")
	return nil
}

func batchInsertUsers(db *sql.DB, userBatch []Account) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}

	insertQuery := `INSERT INTO user_accounts (full_name, years_old) VALUES ($1, $2)`
	for _, user := range userBatch {
		if _, err := tx.Exec(insertQuery, user.FullName, user.YearsOld); err != nil {
			tx.Rollback()
			return fmt.Errorf("insert failure for user %s: %w", user.FullName, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	fmt.Println("All user data inserted successfully!")
	return nil
}

func fetchUsers(db *sql.DB, ageFilter int, resultsLimit, resultsOffset int) ([]Account, error) {
	var rows *sql.Rows
	var err error

	if ageFilter > 0 {
		query := `SELECT id, full_name, years_old FROM user_accounts WHERE years_old = $1 LIMIT $2 OFFSET $3`
		rows, err = db.Query(query, ageFilter, resultsLimit, resultsOffset)
	} else {
		query := `SELECT id, full_name, years_old FROM user_accounts LIMIT $1 OFFSET $2`
		rows, err = db.Query(query, resultsLimit, resultsOffset)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve users: %w", err)
	}
	defer rows.Close()

	var userList []Account
	for rows.Next() {
		var acc Account
		if err := rows.Scan(&acc.ID, &acc.FullName, &acc.YearsOld); err != nil {
			return nil, err
		}
		userList = append(userList, acc)
	}

	return userList, nil
}

func modifyUser(db *sql.DB, id int, newName string, newAge int) error {
	query := `UPDATE user_accounts SET full_name = $1, years_old = $2 WHERE id = $3`
	_, err := db.Exec(query, newName, newAge, id)
	if err != nil {
		return fmt.Errorf("unable to update user: %w", err)
	}

	fmt.Printf("User with ID %d has been updated!\n", id)
	return nil
}

func removeUser(db *sql.DB, id int) error {
	query := `DELETE FROM user_accounts WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("unable to delete user: %w", err)
	}

	fmt.Printf("User with ID %d has been deleted!\n", id)
	return nil
}

type Account struct {
	ID       int
	FullName string
	YearsOld int
}

func main() {
	// Establish connection to the DB
	db, err := establishConnection()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close()

	// Initialize the table
	if err := initializeTable(db); err != nil {
		log.Fatalf("Table creation failed: %v", err)
	}

	users := []Account{
		{FullName: "Anuar Daniyar", YearsOld: 22},
		{FullName: "Aidar Nurkin", YearsOld: 20},
		{FullName: "Dosik Dosik", YearsOld: 21},
	}

	if err := batchInsertUsers(db, users); err != nil {
		log.Fatalf("User insertion failed: %v", err)
	}

	fetchedUsers, err := fetchUsers(db, 0, 2, 0)
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	fmt.Println("Fetched users:", fetchedUsers)

	if err := modifyUser(db, 8, "Sigma Sigma", 32); err != nil {
		log.Fatalf("Update failed: %v", err)
	}

	if err := removeUser(db, 9); err != nil {
		log.Fatalf("Deletion failed: %v", err)
	}
}
