// The *pgx.Conn returned by pgx.Connect() represents a single connection and is not concurrency safe. This is entirely appropriate for a simple command line example such as above. However, for many uses, such as a web application server, concurrency is required. To use a connection pool replace the import github.com/jackc/pgx/v4 with github.com/jackc/pgx/v4/pgxpool and connect with pgxpool.Connect() instead of pgx.Connect().

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	// Loading environment variables
	// dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	// Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)
	dbpool, err := pgxpool.Connect(context.Background(), dbURI)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	var greeting string
	err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
