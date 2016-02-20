package controllers

import (
	"net/http"
	"github.com/lstern/psilibrary/server/models"
	"github.com/lstern/psilibrary/server/repositories"
)

var repository repositories.CategoryRepository

// CategoryUpdate rota teste
func CategoryUpdate(w http.ResponseWriter, r *http.Request) {
	category := new(models.Category)
	f := func(o interface{})(error) { 
		category, _ := o.(models.Category);
		return repository.UpdateCategory(&category, nil, repository) 
	}
	 
	GenericUpdate(category, r, w, f)
}


// CategoryShow TodoShow rota teste
func CategoryShow(w http.ResponseWriter, r *http.Request) {
	idVar := "ID"
	call := func(v int)(interface{}, error){
		return repository.GetCategoryById(v, nil)
	}

	GenericGetByID(w, r, idVar, call)
}


// CategoryCreate TodoCreate rota teste
func CategoryCreate(w http.ResponseWriter, r *http.Request) {
	category := new(models.Category)
	f := func(o interface{})(error) { 
		category, _ := o.(models.Category);
		_, err := repository.CreateCategory(&category, nil, repository)
		return err 
	}

	GenericUpdate(category, r, w, f)
}

// CategoryIndex rota teste
func CategoryIndex(w http.ResponseWriter, r *http.Request) {
	call := func()(interface{}, error){
		return repository.ListCategories(nil)
	}

	GenericList(w, r, call)
}
