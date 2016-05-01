package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/lstern/psilibrary/server/models"
	"github.com/lstern/psilibrary/server/repositories"
)

var entryTypeRepo repositories.EntryTypeRepository

// EntryTypeUpdate rota teste
func EntryTypeUpdate(w http.ResponseWriter, r *http.Request) {
	f := func(body []byte)(error) { 
		entry := new(models.EntryType)

		if err := json.Unmarshal(body, entry); err != nil {
			w.WriteHeader(422) // unprocessable entity
		}

		return entryTypeRepo.Update(entry) 
	}
	 
	GenericUpdate( r, w, f)
}


// EntryTypeShow TodoShow rota teste
func EntryTypeShow(w http.ResponseWriter, r *http.Request) {
	idVar := "entrytypeId"
	call := func(v int)(interface{}, error){
		return entryTypeRepo.GetById(v)
	}

	GenericGetByID(w, r, idVar, call)
}


// EntryTypeCreate TodoCreate rota teste
func EntryTypeCreate(w http.ResponseWriter, r *http.Request) {
	f := func(body []byte)(error) { 
		entry := new(models.EntryType)

		if err := json.Unmarshal(body, entry); err != nil {
			w.WriteHeader(422) // unprocessable entity
		}

		_, err := entryTypeRepo.Create(entry)
		return err 
	}

	GenericUpdate(r, w, f)
}

// EntryTypeIndex rota teste
func EntryTypeIndex(w http.ResponseWriter, r *http.Request) {
	call := func()(interface{}, error){
		return entryTypeRepo.List()
	}

	GenericList(w, r, call)
}
