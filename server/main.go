package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()
	test()

	log.Fatal(http.ListenAndServe(":8080", router))
}
