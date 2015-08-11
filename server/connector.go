package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func test(){
   	log.Printf("acessando banco")
	db, err := sql.Open("mysql",
		"root:pandora@tcp(127.0.0.1:3306)/psilibrary")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
	// do something here
	} else {
    	log.Printf("funcionou")
	}

	defer db.Close()	
}
