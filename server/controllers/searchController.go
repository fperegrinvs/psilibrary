package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"github.com/lstern/psilibrary/server/models"
	"github.com/lstern/psilibrary/server/repositories"
)

func Search(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	query := new(models.SearchQuery)

	if err := json.Unmarshal(body, &query); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	addCors(w, r)

	repo := repositories.MakeSearchRepository();
	result, err := repo.Search(query)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}

	return
}
