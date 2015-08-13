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
		"EntryTypeUpdate",
		[]string{"POST"},
		"/entrytype/update",
		EntryTypeUpdate,
	},
	Route{
		"EntryTypeUpdateOptions",
		[]string{"OPTIONS"},
		"/entrytype/update",
		OptionsHandler,
	},
	Route{
		"TodoCreate",
		[]string{"POST"},
		"/todos",
		TodoCreate,
	},
	Route{
		"EntryTypeShow",
		[]string{"GET", "OPTIONS"},
		"/entrytype/{entrytypeId}",
		EntryTypeShow,
	},
}