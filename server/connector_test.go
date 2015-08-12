package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

func TestMySqlConnection(t *testing.T) {
   	log.Printf("acessando banco!")
	db, err := sql.Open("mysql",
		"root:pandora@tcp(127.0.0.1:3306)/psilibrary")
	
	defer db.Close()
	
	if err != nil {
		t.Error("Erro ao definir conexão com o banco")
	}

	err = db.Ping()
	if err != nil {
		t.Error("Erro ao tentar conectar ao banco")
	}	
}
