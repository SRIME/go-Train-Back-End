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

	ExecuteSelectQuery(dbPool)
	ExecuteFunction(dbPool)
	log.Println("stopping program")
}

func ExecuteSelectQuery(dbPool *pgxpool.Pool) {
	log.Println("starting execution of select query")
	//execute the query and get result rows
	rows, err := dbPool.Query(context.Background(), "select * from accounts;")
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
		id := values[0].(int32)
		name := values[4].(string)
		email := values[5].(string)
		phone := values[6].(string)
		log.Println("[id:", id, ", Name:", name, ", Email:", email, ", Phone:", phone, "]")
	}

}

func ExecuteFunction(dbPool *pgxpool.Pool) {
	log.Println("starting execution of database function and iterating through a single selected row")
	// id can be taken as a user input
	// for the time being, let's hard code it
	id := 1
	var Name string
	var Email string
	var Phone string
	//execute the query and get result rows
	rows, err := dbPool.Query(context.Background(), "select * from accounts where id='1'")
	log.Println("input id: ", id)
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
		Name = values[4].(string)
		Email = values[5].(string)
		Phone = values[6].(string)

		log.Println("[Name:", Name, ", Email:", Email, ", Phone:", Phone, "]")
	}

}
