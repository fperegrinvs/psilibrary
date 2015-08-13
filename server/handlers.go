package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"psilibrary/server/models"
	"psilibrary/server/repositories"
	"github.com/gorilla/mux"
)

// Index Página principal
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func addCors(w http.ResponseWriter, r *http.Request){
    if origin := r.Header.Get("Origin"); origin != "" {
        w.Header().Set("Access-Control-Allow-Origin", origin)
    } else {
    	w.Header().Set("Access-Control-Allow-Origin", "*")
    }
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Requested-With")
    w.Header().Set("Access-Control-Allow-Credentials", "true")	
}

// TodoIndex rota teste
func EntryTypeIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	addCors(w, r)

	w.WriteHeader(http.StatusOK)
	list, _ := repositories.ListEntryTypes()
	if err := json.NewEncoder(w).Encode(list); err != nil {
		panic(err)
	}
}

func OptionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	addCors(w, r)

	w.WriteHeader(http.StatusOK)
}


// TodoShow rota teste
func EntryTypeShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var entryTypeID int
	var err error
	if entryTypeID, err = strconv.Atoi(vars["entrytypeId"]); err != nil {
		panic(err)
	}
	entryType, err := repositories.GetEntryTypeById(entryTypeID)
	if entryType.ID > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		addCors(w, r)
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(entryType); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

// TodoCreate rota teste
func EntryTypeUpdate(w http.ResponseWriter, r *http.Request) {
	entryType := new(models.EntryType)
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &entryType); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	addCors(w, r)
	err = repositories.UpdateEntryType(entryType)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}

// TodoCreate rota teste
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}