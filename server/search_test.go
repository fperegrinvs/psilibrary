package main_test

import (
	"strconv"
	 "testing"
	 "github.com/lstern/psilibrary/server"	
	 "github.com/lstern/psilibrary/server/conf"
	 "github.com/lstern/psilibrary/server/models"
	 "github.com/lstern/psilibrary/server/repositories"
 )

var (
	repo repositories.SearchRepository
)


func init(){
   repo = repositories.MakeSearchRepository()
}

// check is routes are ok
func Test_SearchRoute(t *testing.T){
	router := main.NewRouter()

	if router.Get("Search") == nil {
		t.Error("rota de busca não está registrada")
		return
	}
}

func Test_Process_NullQuery_Should_Give_Errror(t *testing.T){
	_, err := repo.ProcessInput(nil)

	if err == nil  {
		t.Error("Erro ao processar query nula")
		return
	}
}

func Test_Default_Values(t *testing.T){
	query := new(models.SearchQuery)
	query.Query = "batman"

	query, _ = repo.ProcessInput(query)

	if query == nil {
		t.Error("Erro ao processar query")
		return
	}

	if query.PageSize != conf.PageSize  {
		t.Error("Erro ao definir valor padrão para o tamanho da página")
		return
	}

	if query.Page != 1  {
		t.Error("Erro ao definir valor padrão para a página")
		return
	}

	if query.Order != models.Order_Date {
		t.Error("Erro ao definir a ordenação padrão")
		return
	}

}

func Test_Allow_Empty_Query(t *testing.T){
	query := new(models.SearchQuery)
	query.Query = ""

	_, err := repo.ProcessInput(query)

	if err != nil  {
		t.Error("Não deveria ocorrer um erro ao processar uma query vazia")
		return
	}
}

func Test_Process_Result_Navigation_Page_1(t *testing.T){
	query := new(models.SearchQuery)
	query.Query = ""
	query.Page = 1
	query.PageSize = 10

	results := 
		[]*models.Entry{
			&models.Entry{EntryId:1}, 
			&models.Entry{EntryId:2}, 
			&models.Entry{EntryId:3}}

	response, err := repo.ProcessResultsNavigation(query, results, 3)

	if err != nil || response == nil {
		t.Error("Erro ao processar resultados ", err)
		return
	}

	if response.Navigation.TotalPages != 1 {
		t.Error("Total de páginas deve ser 1 : " + strconv.Itoa(response.Navigation.TotalPages))
		return
	}

	if response.Navigation.PageStart != 1 {
		t.Error("Inicio da página deve ser 1")
		return 
	}

	if response.Navigation.PageEnd != 3 {
		t.Error("Fim da página deve ser 3")
		return
	}

	if response.Navigation.CurentPage != 1 {
		t.Error("Página atual deve ser 1")
		return
	}

	if response.Navigation.TotalCount != 3 {
		t.Error("Quantidade total de registros deve ser 3")
		return
	}
}


func Test_Process_Result_Navigation_Page_2(t *testing.T){
	query := new(models.SearchQuery)
	query.Query = ""
	query.Page = 2
	query.PageSize = 1

	results := 
		[]*models.Entry{
			&models.Entry{EntryId:1}, 
			&models.Entry{EntryId:2}, 
			&models.Entry{EntryId:3}}

	response, err := repo.ProcessResultsNavigation(query, results, 3)

	if err != nil || response == nil {
		t.Error("Erro ao processar resultados ", err)
		return
	}

	if response.Navigation.TotalPages != 3 {
		t.Error("Total de páginas deve ser 3 : " + strconv.Itoa(response.Navigation.TotalPages))
		return
	}

	if response.Navigation.PageStart != 2 {
		t.Error("Inicio da página deve ser 2")
		return 
	}

	if response.Navigation.PageEnd != 2 {
		t.Error("Fim da página deve ser 2")
		return
	}

	if response.Navigation.CurentPage != 2 {
		t.Error("Página atual deve ser 2")
		return
	}

	if response.Navigation.TotalCount != 3 {
		t.Error("Quantidade total de registros deve ser 3")
		return
	}
}


func Test_Process_Facets_Category_No_Filter(t *testing.T) {
	results := []*models.Entry{}

	query_result := new(models.SearchResults)
	query_result.Results = results

	response, err := repo.ProcessFacets(query_result)

	if err != nil || response == nil {
		t.Error("Erro ao processar facets ", err)
		return
	}

	if len(response.Facets) == 0 {
		t.Error("Nenhuma facet encontrado")
	}
}

func Test_Process_Facets_Category_Filter(t *testing.T) {
	results := []*models.Entry{}

	query_result := new(models.SearchResults)
	query_result.Results = results

	filters := make(map[string][]string)
	filters["category"] = []string{"2"}
	query_result.Query.Filters = filters


	response, err := repo.ProcessFacets(query_result)

	if err != nil || response == nil {
		t.Error("Erro ao processar facets ", err)
		return
	}

	if len(response.Facets) != 1 {
		t.Error("Um facet deveria ser encontrado")
		return
	}

	if response.Facets[0].Id != "category" {
		t.Error("Facet retornado deveria ser o de categoria")
		return
	}

	if len(response.Facets[0].Options) != 1 {
		t.Error("Facet de categoria deveria ter apenas uma opçào: " + strconv.Itoa( len(response.Facets[0].Options)))
		return
	}

}

func Test_Execute_Search_Empty(t *testing.T) {
	query := new(models.SearchQuery)
	query.Query = ""
	query.Page = 2
	query.PageSize = 1
	response, err := repo.ExecuteSearch(query)

	if err != nil {
		t.Error("Erro ao executar query: ", err)
		return
	}

	if response == nil || len(response) == 0 {
		t.Error("Nenhum resultado retornado", err)
		return
	}
}

func Test_Execute_Search_Filter_NoQuery(t *testing.T) {
	query := new(models.SearchQuery)
	query.Query = ""
	query.Page = 2
	query.PageSize = 1
	query.Filters  = make(map[string][]string)
	query.Filters["category"] = []string{"2"}
	
	response, err := repo.ExecuteSearch(query)

	if err != nil {
		t.Error("Erro ao executar query: ", err)
		return
	}

	if response == nil || len(response) == 0 {
		t.Error("Nenhum resultado retornado", err)
		return
	}
}

func Test_Execute_Search_Query_NoFilter(t *testing.T) {
	query := new(models.SearchQuery)
	query.Query = "abstract"
	query.Page = 2
	query.PageSize = 1
	response, err := repo.ExecuteSearch(query)

	if err != nil {
		t.Error("Erro ao executar query: ", err)
		return
	}

	if response == nil || len(response) == 0 {
		t.Error("Nenhum resultado retornado", err)
		return
	}
}

func Test_Execute_Search_Filter_and_Query(t *testing.T) {
	query := new(models.SearchQuery)
	query.Query = "abstract"
	query.Page = 2
	query.PageSize = 1
	query.Filters  = make(map[string][]string)
	query.Filters["category"] = []string{"2"}
	
	response, err := repo.ExecuteSearch(query)

	if err != nil {
		t.Error("Erro ao executar query: ", err)
		return
	}

	if response == nil || len(response) == 0 {
		t.Error("Nenhum resultado retornado", err)
		return
	}
}

//func Test_ExecuteSearch_No_Filter(t *testing.T) {
