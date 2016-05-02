package repositories

import (
	"errors"
	"math"
	"strconv"
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
	query, err := s.ProcessInput(query)

	if err != nil {
		return nil, err
	}

	results, total, err := s.ExecuteSearch(query)

	if err != nil {
		return nil, err
	}

	response, err := s.ProcessResultsNavigation(query, results, total)

	if err != nil {
		return nil, err
	}

	response, err = s.ProcessFacets(response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s SearchRepository) ExecuteSearch(query *models.SearchQuery) ([]models.Entry, int, error){
	category := query.Filters["category"]

	db, err := openSql(s.DB)

	if err != nil {
		return nil, 0, err
	}	

	defer db.Close()

	result := []models.Entry{}
	start := (query.Page - 1) * query.PageSize 
	var count int
	if query.Query == "" {
		if category == nil {
			err = db.Select(&result, "SELECT * FROM entry limit ?,?", start, query.PageSize)
			
			if err == nil {
				err = db.Get(&count, "SELECT count(*) FROM entry")
			}
		} else {
			catid,_ := strconv.Atoi(category[0])
			err = db.Select(&result, "SELECT e.* FROM entry e inner join categoryentry ce on ce.EntryId = e.EntryId where ce.CategoryId = ? limit ?,?", catid,  start, query.PageSize)

			if err == nil {
				err = db.Get(&count, "SELECT count(*) FROM entry e inner join categoryentry ce on ce.EntryId = e.EntryId where ce.CategoryId = ?", catid)
			}
		}
	} else {
		if category == nil {
			err = db.Select(&result, "call search(?,?,?)", query.Query, start, query.PageSize)

			if err == nil {
				err = db.Get(&count, "call search_count(?)", query.Query)
			}

		} else {
			catid, _ := strconv.Atoi(category[0])
			err = db.Select(&result, "call SearchByCategory(?,?,?,?)", query.Query, catid, start, query.PageSize)

			if err == nil {
				err = db.Get(&count, "call SearchByCategory_count(?,?)", query.Query, catid)
			}
		}
	}

	if err != nil {
		return nil, 0, err
	}

	return result, count, nil
}

func (s SearchRepository) ProcessResultsNavigation(query *models.SearchQuery, results []models.Entry, total int) (*models.SearchResults, error){
	response := new(models.SearchResults)

	response.Navigation.TotalPages =  int(math.Ceil(float64(total) / float64(query.PageSize)))
	response.Navigation.PageStart = 1 + query.PageSize * (query.Page - 1)
	response.Navigation.PageEnd = int(math.Min(float64(response.Navigation.PageStart + query.PageSize - 1), float64(total)))
	response.Navigation.CurentPage = query.Page
	response.Navigation.TotalCount = total
	response.Query = *query

	response.Results = results
	return response, nil
}

func (s SearchRepository) ProcessFacets(results *models.SearchResults) (*models.SearchResults, error){
	category := results.Query.Filters["category"]

	cat_repo := MakeCategoryRepository(nil, nil)
	cats := []*models.Category{}
	var filtered bool

	if category == nil {
		filtered = false
		cats, _ = cat_repo.List();
	} else {
		filtered = true
   		var ids = []int64{}

   	    for _, i := range category {
	        j, err := strconv.Atoi(i)
	        if err != nil {
	            panic(err)
	        }
	        ids = append(ids, int64(j))
	    }

		cats_p, _ := cat_repo.GetCategoriesByIdList(ids);
		for _, cat := range cats_p {
			cats = append(cats, &cat)
		}
	}

	var options = []models.FacetOption{}
	for _, cat := range cats {
		option := models.FacetOption{}
		option.Id = cat.ID
		option.Name = cat.Name
		option.IsSelected = filtered
		options = append(options, option)
	}

	facet := models.Facet{}
	facet.Id = "category"
	facet.Name = "Categorias"
	facet.Options = options
	facet.IsSelected = filtered

	results.Facets = []models.Facet{facet}

	return results, nil
}
