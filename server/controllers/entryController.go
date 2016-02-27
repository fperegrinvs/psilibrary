package controllers

import (
	"net/http"
	"github.com/lstern/psilibrary/server/models"
	"github.com/lstern/psilibrary/server/repositories"
)

var entryRepo = repositories.MakeEntryRepository(nil)

// CategoryUpdate rota teste
func EntryUpdate(w http.ResponseWriter, r *http.Request) {
	entry := new(models.Entry)
	f := func(o interface{})(error) { 
		entry, _ := o.(models.Entry);
		return entryRepo.Update(&entry) 
	}
	 
	GenericUpdate(entry, r, w, f)
}


// CategoryShow TodoShow rota teste
func EntryUShow(w http.ResponseWriter, r *http.Request) {
	idVar := "ID"
	call := func(v int)(interface{}, error){
		return entryRepo.GetById(v)
	}

	GenericGetByID(w, r, idVar, call)
}


// CategoryCreate TodoCreate rota teste
func EntryCreate(w http.ResponseWriter, r *http.Request) {
	entry := new(models.Entry)
	f := func(o interface{})(error) { 
		entry, _ := o.(models.Entry);
		_, err := entryRepo.Create(&entry)
		return err 
	}

	GenericUpdate(entry, r, w, f)
}

// CategoryIndex rota teste
func EntryIndex(w http.ResponseWriter, r *http.Request) {
	call := func()(interface{}, error){
		return entryRepo.List()
	}

	GenericList(w, r, call)
}
