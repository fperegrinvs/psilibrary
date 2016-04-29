package main_test

import (
	 "testing"
	 "github.com/lstern/psilibrary/server"
 )

// check is routes are ok
func Test_SearchRoute(t *testing.T){
	router := main.NewRouter()

	if router.Get("Search") == nil {
		t.Error("rota de busca não está registrada")
	}
}
