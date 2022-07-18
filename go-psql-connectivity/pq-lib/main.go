package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "2fb4cac9"
	dbName   = "efarm"
)

func main() {
	dbURL := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to psql")

	// insert into table:

	// Just add the record into the table
	// insertValue(db)
	//insert and return the value of the newly added record from table
	// insertAndReturnValueFromDB(db)

	// Update the Value in Table
	// updateValue(db)
	// Update the value and return the new values from table
	// updateValuePrintIt(db)

	// Delete the Row in table
	deleteValue(db)

	// Print Table
	// printTable(db)
}

// func insertValue(db *sql.DB) {
// 	Name := "Naruto"
// 	Email := "naruto@email.com"
// 	Phone := 6656875497
// 	insertStatement := `insert into demoacc(name, email, phone)
// 						values ($1, $2, $3)`
// 	db.Exec(insertStatement, Name, Email, Phone)
// }

// func insertAndReturnValueFromDB(db *sql.DB) {
// 	insertStatement := `insert into demoacc(name , email, phone) values ($1, $2, $3) returning id`

// 	Name := "Kevin"
// 	Email := "kev@mail.com"
// 	Phone := 734519871
// 	id := 0
// 	err := db.QueryRow(insertStatement, Name, Email, Phone).Scan(&id)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Added New Record: ", id)
// }

// func updateValue(db *sql.DB) {
// 	sqlStatement := `UPDATE demoacc SET email=$2 where name=$1`
// 	db.Exec(sqlStatement, "Kevin", "kev1@mail.com")
// }

// func updateValuePrintIt(db *sql.DB) {
// 	var Name string
// 	var Email string
// 	var Phone string
// 	var id int
// 	sqlStatement := `UPDATE demoacc SET name=$1 WHERE name = $2 RETURNING name,email,phone,id;`
// 	err := db.QueryRow(sqlStatement, "Hari V", "Hari").Scan(&Name, &Email, &Phone, &id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(Name, Email, Phone, id)

// }

func deleteValue(db *sql.DB) {
	sqlStatement := `DELETE from demoacc where id = $1`
	res, err := db.Exec(sqlStatement, 8)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	if count != 1 {
		fmt.Println("Error while deleting your row")
	}
	if count == 1 {
		fmt.Println("Successfully deleted!")
	}
}

// func printTable(db *sql.DB) {
// 	stmt, err := db.Prepare("select *from ?")
// 	if err != nil {
// 		panic(err)
// 	}
// 	stmt.Exec("demoacc")

// }
