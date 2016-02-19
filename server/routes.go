package main

import ("net/http" 
	"psilibrary/server/controllers"
)


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
		controllers.Index,
	},
	Route{
		"EntryTypeIndex",
		[]string{"GET","OPTIONS"},
		"/entrytype",
		controllers.EntryTypeIndex,
	},
	Route{
		"EntryTypeUpdate",
		[]string{"POST"},
		"/entrytype/update",
		controllers.EntryTypeUpdate,
	},
	Route{
		"EntryTypeUpdateOptions",
		[]string{"OPTIONS"},
		"/entrytype/{value}",
		controllers.OptionsHandler,
	},
	Route{
		"EntryTypeCreate",
		[]string{"POST"},
		"/entrytype/create",
		controllers.EntryTypeCreate,
	},
	Route{
		"EntryTypeShow",
		[]string{"GET"},
		"/entrytype/{entrytypeId}",
		controllers.EntryTypeShow,
	},

	// Category

	Route{
		"CategoryIndex",
		[]string{"GET","OPTIONS"},
		"/category",
		controllers.CategoryIndex,
	},
	Route{
		"CategoryUpdate",
		[]string{"POST"},
		"/category/update",
		controllers.CategoryUpdate,
	},
	Route{
		"CategoryUpdateOptions",
		[]string{"OPTIONS"},
		"/category/{value}",
		controllers.OptionsHandler,
	},
	Route{
		"CategoryCreate",
		[]string{"POST"},
		"/category/create",
		controllers.CategoryCreate,
	},
	Route{
		"CategoryGet",
		[]string{"GET"},
		"/category/{CategoryId}",
		controllers.CategoryShow,
	},
}