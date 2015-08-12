package repositories

import (
	"psilibrary/server/models"
	"psilibrary/server/conf"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ListEntryTypes() []*models.EntryType {
	var entries []*models.EntryType

	db, err := sql.Open(conf.Db, conf.Conn)	
	defer db.Close()

	rows, err := db.Query("SELECT EntryTypeId, Name FROM entrytype")

	for rows.Next() {
	    e := new(models.EntryType)
	    if err := rows.Scan(&e.ID, &e.Name); err != nil { }
	    entries = append(entries, e)
	}

	if err != nil {
	}

	return entries
}
