package models

type (
	// EntryType é o tipo de publicação de um registro
	Category struct {
		ID     int 			 `json:"id" db:"CategoryId"`
		Name   string        `json:"name" db:"Name"`
		ParentId int         `json:"parentId" db:"ParentId"`
	}
)