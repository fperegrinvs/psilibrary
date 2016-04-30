package controllers

import (
	"errors"
	"fmt"
	"reflect"
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
func EntryShow(w http.ResponseWriter, r *http.Request) {
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
		entry, ok := o.(*models.Entry);

		if !ok  {
			print("Erro ao converter tipo:\n")
			fmt.Println(reflect.TypeOf(o))
			return errors.New("Erro ao converter tipo");
		}

		_, err := entryRepo.Create(entry)
		

		return err 
	}

	GenericUpdate(entry, r, w, f)
}

// CategoryIndex rota teste
func EntryIndex(w http.ResponseWriter, r *http.Request) {
	call := func()(interface{}, error){
		return entryRepo.List(0,99999)
	}

	GenericList(w, r, call)
}
