package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/lstern/psilibrary/server/models"
	"github.com/lstern/psilibrary/server/repositories"
)

var catRepository = repositories.MakeCategoryRepository(nil, nil)

// CategoryUpdate rota teste
func CategoryUpdate(w http.ResponseWriter, r *http.Request) {

	f := func(body []byte)(error) { 
		category := new(models.Category)
		if err := json.Unmarshal(body, category); err != nil {
			w.WriteHeader(422) // unprocessable entity
		}

		return catRepository.Update(category) 
	}
	 
	GenericUpdate(r, w, f)
}


// CategoryShow TodoShow rota teste
func CategoryShow(w http.ResponseWriter, r *http.Request) {
	idVar := "CategoryId"
	call := func(v int)(interface{}, error){
		return catRepository.GetById(v)
	}

	GenericGetByID(w, r, idVar, call)
}


// CategoryCreate TodoCreate rota teste
func CategoryCreate(w http.ResponseWriter, r *http.Request) {
	f := func(body []byte)(error) { 
		category := new(models.Category)
		
		if err := json.Unmarshal(body, category); err != nil {
			w.WriteHeader(422) // unprocessable entity
		}

		_, err := catRepository.Create(category)
		return err 
	}

	GenericUpdate(r, w, f)
}

// CategoryIndex rota teste
func CategoryIndex(w http.ResponseWriter, r *http.Request) {
	call := func()(interface{}, error){
		return catRepository.List()
	}

	GenericList(w, r, call)
}
