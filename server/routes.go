package main

import "net/http"

// Route rota de acesso ao serviço REST
type Route struct {
	Name        string
	Methods     []string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes lista de rotas do REST
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		[]string{"GET"},
		"/",
		Index,
	},
	Route{
		"EntryTypeIndex",
		[]string{"GET","OPTIONS"},
		"/entrytype",
		EntryTypeIndex,
	},
	Route{
		"TodoCreate",
		[]string{"POST"},
		"/todos",
		TodoCreate,
	},
	Route{
		"TodoShow",
		[]string{"GET"},
		"/todos/{todoId}",
		TodoShow,
	},
}