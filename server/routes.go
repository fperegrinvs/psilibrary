package main

import ("net/http" 
	"github.com/lstern/psilibrary/server/controllers"
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

	//Entry
	Route{
		"EntryIndex",
		[]string{"GET","OPTIONS"},
		"/entry",
		controllers.EntryIndex,
	},
	Route{
		"EntryUpdate",
		[]string{"POST"},
		"/entry/update",
		controllers.EntryUpdate,
	},
	Route{
		"EntryUpdateOptions",
		[]string{"OPTIONS"},
		"/entry/{value}",
		controllers.OptionsHandler,
	},
	Route{
		"EntryCreate",
		[]string{"POST"},
		"/entry/create",
		controllers.EntryCreate,
	},
	Route{
		"EntryGet",
		[]string{"GET"},
		"/entry/{CategoryId}",
		controllers.EntryShow,
	},
	Route{
		"FacebookCallback",
		[]string{"GET"},
		"/auth/facebook/callback",
		controllers.FacebookCallback,
	},
	Route{
		"SearchOptions",
		[]string{"OPTIONS"},
		"/search",
		controllers.OptionsHandler,
	},

	Route{
		"Search",
		[]string{"POST"},
		"/search",
		controllers.Search,
	},
	Route{
		"MedlineOptions",
		[]string{"OPTIONS"},
		"/medline",
		controllers.OptionsHandler,
	},

	Route{
		"Medline",
		[]string{"POST"},
		"/medline",
		controllers.ImportMedline,
	},
	Route{
		"UserGet",
		[]string{"GET", "OPTIONS"},
		"/user/{id}",
		controllers.CheckUser,
	},
}