package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
    "github.com/stretchr/gomniauth"  
	"github.com/stretchr/gomniauth/providers/facebook"
)

var router *mux.Router;


func Setup() {
	router = NewRouter()
	gomniauth.SetSecurityKey("hPIgpDjBNmZxEMU7leIxm892FhK7HG5cArZUj1lQl7Qmy55izmlfkL3OsAzHnTYZ")
	gomniauth.WithProviders(
		facebook.New("268751096793071", "02b6546587caca852c3e9c117a2d6024", "http://psi-library.azurewebsites.net/auth/facebook/callback"),
	)
}

func main() {
	Setup();
	log.Fatal(http.ListenAndServe(":80", router))
}
