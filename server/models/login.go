package models

import "time"

type (
	// EntryType é o tipo de publicação de um registro
	Token struct {
		ID     string 			 `json:"id" db:"SessionId"`
		Expiration time.Time `json:"expiration" db:"Expiration"`
		IsValid bool         `json:"isValid"`
		User
	}

	Session struct {
		ID string	`json:"id" db:"SessionId"`
		Login  string		 `json:"login" db:"Login"`
		Expiration time.Time `json:"expiration" db:"Expiration"`
	}
)

