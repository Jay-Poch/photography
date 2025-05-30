package main

import (
	"database/sql"
	"fmt"
	"log"
)

//const username = "root"
//const password = "REMOVED"
//const address = "192.168.178.108:3306"

const dsn = username + ":" + password + "@tcp(" + address + ")/mydb"

func connection(dsn string) *sql.DB {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	// Check the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}
	fmt.Println("Connected to MySQL!")
	return db
}

func selecty(querry string) {
	db := connection(dsn)
	rows, err := db.Query(querry)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Println(rows)
	for rows.Next() {
		var ID string
		var Location string
		var Info string
		if err := rows.Scan(&ID, &Location, &Info); err != nil {
			log.Fatal(err)
		}
		//respone = image{Id: ID, Location: Location, Info: map[string]string{"test": "test"}}
		fmt.Printf("ID: %s, Location: %s, Info: %s\n", ID, Location, Info)
	}

	defer db.Close()

}
