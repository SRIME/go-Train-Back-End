package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	log.Println("starting program")
	// get the database connection URL.
	// usually, this is taken as an environment variable as in below commented out code
	// databaseUrl = os.Getenv("DATABASE_URL")
	// for the time being, let's hard code it as follows. change the values as needed.

	databaseUrl := "postgres://postgres:2fb4cac9@localhost:5432/efarm"
	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	//to close DB pool
	defer dbPool.Close()

	// ExecuteFunction(dbPool)
	ExecuteDeleteQuery(dbPool)
	ExecuteSelectQuery(dbPool)
	log.Println("stopping program")
}

func ExecuteSelectQuery(dbPool *pgxpool.Pool) {
	log.Println("starting execution of select query")
	//execute the query and get result rows
	rows, err := dbPool.Query(context.Background(), "select * from demoacc")
	if err != nil {
		log.Fatal("error while executing query")
	}

	log.Println("result:")
	//iterate through the rows
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}
		//convert DB types to Go types
		// id := values[0].(int32)
		name := values[0].(string)
		email := values[1].(string)
		phone := values[2].(string)
		log.Println("[id:", ", Name:", name, ", Email:", email, ", Phone:", phone, "]")
	}

}

func ExecuteFunction(dbPool *pgxpool.Pool) {
	log.Println("starting execution of database function and iterating through a single selected row")
	// id can be taken as a user input
	// for the time being, let's hard code it

	var Name string
	var Email string
	var Phone string
	Name = "karthick"
	// execute the query and get result row
	// $1 is like a placeholder for variable defined after that
	// rows, err := dbPool.QueryRow("select name, email from demoacc where name=$1", Name).Scan(&Name, &Email)
	rows, err := dbPool.Query(context.Background(), "select phone, email from demoacc where name=$1;", Name)
	if err != nil {
		log.Fatal("error while executing query")
	}

	log.Println("result:")
	//iterate through the rows
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}

		//convert DB types to Go types
		Name = values[0].(string)
		Email = values[1].(string)
		Phone = values[2].(string)

		log.Println("[Name:", Name, ", Email:", Email, ", Phone:", Phone, "]")
	}

}

func ExecuteDeleteQuery(dbPool *pgxpool.Pool) {
	// Exec is used to execute queries that do not return anything like deleting/creating

	fmt.Println("Executing Delete Query")
	Name := "karthick"
	deleter, err := dbPool.Exec(context.Background(), "delete from demoacc where name=$1", Name)

	if err != nil {
		log.Fatal("Error while executing the delete query")
	}
	if deleter.RowsAffected() != 1 {
		log.Fatal("No rows found to delete")
	}
}
