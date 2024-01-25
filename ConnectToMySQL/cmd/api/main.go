package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type Application struct {
	People []Person
}

type connection struct {
	User string
	PW string
	Hostname string
	Net string
	Database string
}

type Person struct {
	ID        int
	FirstName string
	LastName  string
	Age int
	Address string
}

func main() {
	var app Application
	
	cfg := connection{
		User: "root",
		PW: "0987%^poIU",
		Hostname: "localhost:3306",
		Net: "tcp",
		Database: "people",
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@%v(%v)/%v",cfg.User, cfg.PW, cfg.Net, cfg.Hostname, cfg.Database ))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	reader, err := db.Query(fmt.Sprintf("SELECT * FROM %v", cfg.Database))
	if err != nil{
		log.Fatal(err)
	}

	defer reader.Close()

	for reader.Next() {
		var p Person
		
		err = reader.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Age, &p.Address)
		if err != nil {
			log.Fatal(err)
		}
		app.People = append(app.People, p)
	}
}