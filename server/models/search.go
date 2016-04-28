package models

type (
	SearchQuery struct {
		Query     string 			 `json:"query"`
		Filters   []SearchFilter     `json:"filters"`
		PageSize int			     `json:"pageSize"`
		Page int  					 `json:"page"`
		Order string 				 `json:"order"`
	}

	SearchResults struct {
		Query     SearchQuery 	 `json:"query"`
		Results   []Entry        `json:"results"`
		Navigation Navigation     `json:"navigation"`
		Facets []Facet 			`json:"facets"`
	}

	SearchFilter struct {
		Name	   string	`json:"name"`
		Values 	   []string	`json:"values"`
	}

	Navigation struct {
		Order		string 		`json:"order"`
		CurentPage     int 	 	`json:"currentPage"`
		PageSize	int     	`json:"pageSize"`
		PageStart	int         `json:"pageStart"`
		PageEnd 	int         `json:"pageEnd"`
		TotalPages   int        `json:"totalPages"`
		TotalProducts int 		`json:"totalProducts"`
	}

	Facet struct {
		Id int          	`json:"id"` 
		Name string         `json:"name"` 
		IsSelected bool		`json:"isSelected"`
		Options []FacetOption  `json:"options"`
	}

	FacetOption struct {
		Id int          	`json:"id"` 
		Name string         `json:"name"` 
		IsSelected bool		`json:"isSelected"`
		Count int           `json:"count"`
		Options []FacetOption  `json:"options"`
	}
)
