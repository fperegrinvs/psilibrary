package models

type (
	// EntryType é o tipo de publicação de um registro
	User struct {
		Login  string		 `json:"login" db:"Login"`
		Permissions  string	 `json:"permissions" db:"Permissions"`
		Name   string        `json:"name" db:"Name"`
	}
)