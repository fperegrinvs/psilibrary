package repositories

import (
	"errors"
	"math"
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

func (s SearchRepository) ProcessResultsNavigation(query *models.SearchQuery, results []*models.Entry, total int) (*models.SearchResults, error){
	response := new(models.SearchResults)

	response.Navigation.TotalPages =  int(math.Ceil(float64(total) / float64(query.PageSize)))
	response.Navigation.PageStart = 1 + query.PageSize * (query.Page - 1)
	response.Navigation.PageEnd = response.Navigation.PageStart + int(math.Min(float64(len(results)), float64(query.PageSize))) - 1
	response.Navigation.CurentPage = query.Page
	response.Navigation.TotalCount = total

	response.Results = results


	return response, nil
}

func (s SearchRepository) ProcessCategoryFacet(results *models.SearchResults) (*models.SearchResults, error){
	return results, nil
}
