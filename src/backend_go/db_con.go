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
		var camera_model string
		var lens_used string
		var shutter_speed string
		var aperture string
		var ISO string
		var lighting string
		var focal_point string
		var photo_genre string

		if err := rows.Scan(
			&ID,
			&Location,
			&camera_model,
			&lens_used,
			&shutter_speed,
			&aperture,
			&ISO,
			&lighting,
			&focal_point,
			&photo_genre,
		); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %s, Location: %s, camera_model: %s,lens_used: %s,shutter_speed: %s,aperture: %s,ISO: %s,lighting: %s,focal_point: %s,photo_genre: %s\n",
			ID,
			Location,
			camera_model,
			lens_used,
			shutter_speed,
			aperture,
			ISO,
			lighting,
			focal_point,
			photo_genre,
		)
		//respone := image{Id: ID, Location: Location, Info: map[string]string{"test": "test"}}
	}

	defer db.Close()

}
