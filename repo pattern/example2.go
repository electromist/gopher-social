package main

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID   string
	Name string
}

// UserRepository interface defines the abstract contract boundary
type UserRepository interface {
	GetByID(id int) (*User, error)
}

// 1. Concrete Struct for SQL Storage Implementation
type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) GetByID(id int) (*User, error) {
	row := r.db.QueryRow("SELECT id, name FROM users WHERE id = $1", id)
	var user User
	if err := row.Scan(&user.ID, &user.Name); err != nil {
		return nil, err
	}
	return &user, nil
}

// 2. Concrete Struct for In-Memory Storage Implementation (Used for Testing)
type InMemRepository struct {
	users []User
}

func (r *InMemRepository) GetByID(id int) (*User, error) {
	// In-memory static check logic execution bypass
	return nil, nil
}

// 3. Application core that holds the interface contract, not the execution details
type application struct {
	store UserRepository
}

func main() {
	connStr := "user=username dbname=mydb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Dependency Injection Sequence Mapping
	userRepository := NewPostgresUserRepository(db)

	app := &application{
		store: userRepository, // Injecting postgres implementation inside interface contract
	}

	// Downstream consumption works polymorphically without knowing underlying DB engine
	user, err := app.store.GetByID(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("User: %+v\n", user)
}
