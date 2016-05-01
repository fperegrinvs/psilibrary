package controllers

import (
	"log"
	"net/http"
	"encoding/json"
	//"github.com/lstern/psilibrary/server/models"
	"github.com/lstern/psilibrary/server/repositories"
	"github.com/gorilla/mux"
)

type AuthController struct{
}

func FacebookCallback(w http.ResponseWriter, r *http.Request){
	addCors(w, r)
}

func CheckUser(w http.ResponseWriter, r *http.Request){
	addCors(w, r)

	var repo repositories.UserRepository;

	vars := mux.Vars(r)
	id := vars["id"]

	var err error

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	obj, err := repo.GetById(id)

	if err == nil {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(obj); err != nil {
			log.Printf(err.Error());
			panic(err)
		}

		return
	}

	// If we didn't find it, 404
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}
