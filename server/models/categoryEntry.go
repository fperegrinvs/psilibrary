package models
import "time"

type (
	// EntryType é o tipo de publicação de um registro
	CategoryEntry struct {
		EntryID     int 			 `json:"entryId"`
		CategoryId   string        `json:"categoryId"`
	}
) 