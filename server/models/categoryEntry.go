package models

type (
	// EntryType é o tipo de publicação de um registro
	CategoryEntry struct {
		EntryID     int 			 `json:"entryId"`
		CategoryID   string        `json:"categoryId"`
	}
) 