package controllers

import (
	"net/http"
	"github.com/lstern/psilibrary/server/models"
	"github.com/lstern/psilibrary/server/repositories"
)

var entryTypeRepo repositories.EntryTypeRepository

// EntryTypeUpdate rota teste
func EntryTypeUpdate(w http.ResponseWriter, r *http.Request) {
	entryType := new(models.EntryType)
	f := func(o interface{})(error) { 
		entry, _ := o.(models.EntryType);
		return entryTypeRepo.Update(&entry) 
	}
	 
	GenericUpdate(entryType, r, w, f)
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
	entryType := new(models.EntryType)
	f := func(o interface{})(error) { 
		entry, _ := o.(models.EntryType);
		_, err := entryTypeRepo.Create(&entry)
		return err 
	}

	GenericUpdate(entryType, r, w, f)
}

// EntryTypeIndex rota teste
func EntryTypeIndex(w http.ResponseWriter, r *http.Request) {
	call := func()(interface{}, error){
		return entryTypeRepo.List()
	}

	GenericList(w, r, call)
}
