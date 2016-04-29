package repositories

import (
	"errors"
    "github.com/lstern/psilibrary/server/conf"
	"github.com/lstern/psilibrary/server/models"
)

type SearchRepository struct{
	Repository
}

func MakeSearchRepository() SearchRepository {
	var r SearchRepository
	return r
}

func (s SearchRepository) ProcessInput(query *models.SearchQuery) (*models.SearchQuery, error){
	if (query == nil){
       	return nil, errors.New("Query de busca inexistente")
	}

	if query.PageSize == 0 {
		query.PageSize = conf.PageSize
	}

	if query.Page == 0 {
		query.Page = 1
	}

	if query.Order == "" {
		query.Order = models.Order_Date
	}

	return query, nil
}

func (s SearchRepository) Search(query *models.SearchQuery) (*models.SearchResults, error){
	return nil, nil
}
