package repositories

import (
	"psilibrary/server/models"
	"psilibrary/server/conf"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func AddEntryType(e *models.EntryType) {
	db, err := sql.Open(conf.Db, conf.Conn)	
	defer db.Close()

	_, err = db.Exec("insert into EntryType (Name) values (?)", e.Name)

	if err != nil {}
}

func DeleteEntryType(id int) {
	db, err := sql.Open(conf.Db, conf.Conn)	
	defer db.Close()

	_, err = db.Exec("delete from EntryType where IdEntryType = ?", id)

	if err != nil {}
}


func GetEntryTypeById(id int) *models.EntryType {
	db, err := sql.Open(conf.Db, conf.Conn)	
	defer db.Close()

	rows, err := db.Query("SELECT EntryTypeId, Name FROM entrytype where EntryTypeId = ?", id)

	for rows.Next() {
	    e := new(models.EntryType)
	    if err := rows.Scan(&e.ID, &e.Name); err != nil { }
	    return e
	}

	if err != nil {
	}

	return nil
}

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
