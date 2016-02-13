package models
import "time"

type (
	// EntryType é o tipo de publicação de um registro
	Entry struct {
		ID     int 			 `json:"id"`
		Title   string        `json:"title"`
		Abstract   string        `json:"abstract"`
		Content   string        `json:"content"`
		PublishDate   time.Time        `json:"publishDate"`
		Author   string        `json:"author"`
		Journal   string        `json:"journal"`
		EntryTypeId   int        `json:"entryTypeId"`
	}
)