package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

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


// OptionsHandler handle options
func OptionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	addCors(w, r)

	w.WriteHeader(http.StatusOK)
}

// GenericList is a function to handle listings
func GenericList(w http.ResponseWriter, r *http.Request, call func()(interface{}, error)) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	addCors(w, r)

	w.WriteHeader(http.StatusOK)
	list, err := call()

	if err != nil{
		panic(err)
	}

	if err = json.NewEncoder(w).Encode(list); err != nil {
		panic(err)
	}
}

// GenericGetByID is a functiont to handle get by Id
func GenericGetByID(w http.ResponseWriter, r *http.Request, idVar string, call func(v int)(interface{}, error)) {
	vars := mux.Vars(r)
	var id int
	var err error
	if id, err = strconv.Atoi(vars[idVar]); err != nil {
		panic(err)
	}

	obj, err := call(id)
	if err == nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		addCors(w, r)
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(obj); err != nil {
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


// GenericUpdate is a function to handle updates
func GenericUpdate(obj interface{}, r* http.Request, w http.ResponseWriter, call func(v interface{})(error)){
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &obj); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	addCors(w, r)


	err = call(obj)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}
