package models

type (
	// EntryType é o tipo de publicação de um registro
	EntryType struct {
		ID     int 			 `json:"id"`
		Name   string        `json:"name"`
	}
)