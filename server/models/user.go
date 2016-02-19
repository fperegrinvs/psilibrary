package models

type (
	// EntryType é o tipo de publicação de um registro
	User struct {
		ID     int 			 `json:"id"`
		Login  string		 `json:"login"`
		Password  string	 `json:"password"`
		Name   string        `json:"name"`
	}
)