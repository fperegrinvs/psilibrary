package controllers

import (
	"net/http"
	"psilibrary/server/models"
	"psilibrary/server/repositories"
)

// EntryTypeUpdate rota teste
func EntryTypeUpdate(w http.ResponseWriter, r *http.Request) {
	entryType := new(models.EntryType)
	f := func(o interface{})(error) { 
		entry, _ := o.(models.EntryType);
		return repositories.UpdateEntryType(&entry) 
	}
	 
	GenericUpdate(entryType, r, w, f)
}


// EntryTypeShow TodoShow rota teste
func EntryTypeShow(w http.ResponseWriter, r *http.Request) {
	idVar := "entrytypeId"
	call := func(v int)(interface{}, error){
		return repositories.GetEntryTypeById(v)
	}

	GenericGetByID(w, r, idVar, call)
}


// EntryTypeCreate TodoCreate rota teste
func EntryTypeCreate(w http.ResponseWriter, r *http.Request) {
	entryType := new(models.EntryType)
	f := func(o interface{})(error) { 
		entry, _ := o.(models.EntryType);
		_, err := repositories.CreateEntryType(&entry)
		return err 
	}

	GenericUpdate(entryType, r, w, f)
}

// EntryTypeIndex rota teste
func EntryTypeIndex(w http.ResponseWriter, r *http.Request) {
	call := func()(interface{}, error){
		return repositories.ListEntryTypes()
	}

	GenericList(w, r, call)
}
