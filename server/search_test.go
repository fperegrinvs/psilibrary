package main_test

import (
	 "testing"
	 "github.com/lstern/psilibrary/server"	
	 "github.com/lstern/psilibrary/server/conf"
	 "github.com/lstern/psilibrary/server/models"
	 "github.com/lstern/psilibrary/server/repositories"
 )

// check is routes are ok
func Test_SearchRoute(t *testing.T){
	router := main.NewRouter()

	if router.Get("Search") == nil {
		t.Error("rota de busca não está registrada")
	}
}

func Test_Process_NullQuery_Should_Give_Errror(t *testing.T){
	repo := repositories.MakeSearchRepository()
	_, err := repo.ProcessInput(nil)

	if err == nil  {
		t.Error("Erro ao processar query nula")
	}
}

func Test_Default_Values(t *testing.T){
	query := new(models.SearchQuery)
	query.Query = "batman"

	repo := repositories.MakeSearchRepository()
	query, _ = repo.ProcessInput(query)

	if query == nil {
		t.Error("Erro ao processar query")
	}

	if query.PageSize != conf.PageSize  {
		t.Error("Erro ao definir valor padrão para o tamanho da página")
	}

	if query.Page != 1  {
		t.Error("Erro ao definir valor padrão para a página")
	}

	if query.Order != models.Order_Date {
		t.Error("Erro ao definir a ordenação padrão")
	}

}

func Test_Allow_Empty_Query(t *testing.T){
	query := new(models.SearchQuery)
	query.Query = ""

	repo := repositories.MakeSearchRepository()
	_, err := repo.ProcessInput(query)

	if err != nil  {
		t.Error("Não deveria ocorrer um erro ao processar uma query vazia")
	}
}

//func Test_ExecuteSearch_No_Filter(t *testing.T) {
