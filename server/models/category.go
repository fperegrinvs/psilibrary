package models

type (
	// EntryType é o tipo de publicação de um registro
	Category struct {
		ID     int 			 `json:"id"`
		Name   string        `json:"name"`
		ParentId int         `json:"parentId"`
	}
)