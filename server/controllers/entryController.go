package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/lstern/psilibrary/server/models"
	"github.com/lstern/psilibrary/server/repositories"
)

var entryRepo = repositories.MakeEntryRepository(nil)

// CategoryUpdate rota teste
func EntryUpdate(w http.ResponseWriter, r *http.Request) {
	f := func(body []byte)(error) { 
		entry := new(models.Entry)
		if err := json.Unmarshal(body, entry); err != nil {
			w.WriteHeader(422) // unprocessable entity
		}

		return entryRepo.Update(entry) 
	}
	 
	GenericUpdate(r, w, f)
}


// CategoryShow TodoShow rota teste
func EntryShow(w http.ResponseWriter, r *http.Request) {
	idVar := "CategoryId"
	call := func(v int)(interface{}, error){
		return entryRepo.GetById(v)
	}

	GenericGetByID(w, r, idVar, call)
}


// CategoryCreate TodoCreate rota teste
func EntryCreate(w http.ResponseWriter, r *http.Request) {
	f := func(body []byte)(error) { 
		entry := new(models.Entry)

		if err := json.Unmarshal(body, entry); err != nil {
			w.WriteHeader(422) // unprocessable entity
		}

		_, err := entryRepo.Create(entry)
		
		return err 
	}

	GenericUpdate(r, w, f)
}

// CategoryIndex rota teste
func EntryIndex(w http.ResponseWriter, r *http.Request) {
	call := func()(interface{}, error){
		return entryRepo.List(0,99999)
	}

	GenericList(w, r, call)
}
