package controllers

import (
	"net/http"
	"github.com/lstern/psilibrary/server/models"
	"github.com/lstern/psilibrary/server/repositories"
)

var catRepository repositories.CategoryRepository

// CategoryUpdate rota teste
func CategoryUpdate(w http.ResponseWriter, r *http.Request) {
	category := new(models.Category)
	f := func(o interface{})(error) { 
		category, _ := o.(models.Category);
		return catRepository.Update(&category, catRepository) 
	}
	 
	GenericUpdate(category, r, w, f)
}


// CategoryShow TodoShow rota teste
func CategoryShow(w http.ResponseWriter, r *http.Request) {
	idVar := "ID"
	call := func(v int)(interface{}, error){
		return catRepository.GetById(v)
	}

	GenericGetByID(w, r, idVar, call)
}


// CategoryCreate TodoCreate rota teste
func CategoryCreate(w http.ResponseWriter, r *http.Request) {
	category := new(models.Category)
	f := func(o interface{})(error) { 
		category, _ := o.(models.Category);
		_, err := catRepository.Create(&category, catRepository)
		return err 
	}

	GenericUpdate(category, r, w, f)
}

// CategoryIndex rota teste
func CategoryIndex(w http.ResponseWriter, r *http.Request) {
	call := func()(interface{}, error){
		return catRepository.List()
	}

	GenericList(w, r, call)
}
