package models
import "time"

type (
	// EntryType é o tipo de publicação de um registro
	Entry struct {
		EntryId     int 			 `json:"id" db:"EntryId"`
		Title   string        `json:"title" db:"Title"`
		Abstract   string        `json:"abstract" db:"Abstract"`
		Content   string        `json:"content" db:"Content"`
		PublishDate   time.Time        `json:"publishDate" db:"PublishDate"`
		Author   string        `json:"author" db:"Author"`
		Journal   string        `json:"journal" db:"Journal"`
		EntryType   EntryType        `json:"entryType"`
		EntryTypeId int `db:"EntryTypeId"`
		Categories []Category 	`json:"categories" db:"Categories"`
	}
)