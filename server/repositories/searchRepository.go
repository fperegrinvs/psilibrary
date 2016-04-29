package repositories

import (
	"github.com/lstern/psilibrary/server/models"
)

type SearchRepository struct{
	Repository
}

func MakeSearchRepository() SearchRepository {
	var r SearchRepository
	return r
}

func (s SearchRepository) Search(query *models.SearchQuery) (*models.SearchResults, error){
	return nil, nil
}
