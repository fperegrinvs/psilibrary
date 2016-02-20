package repositories

import (
	"log"
	"github.com/lstern/psilibrary/server/models"
	"github.com/lstern/psilibrary/server/conf"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func CreateEntryType(e *models.EntryType) (int, error) {
	db, err := sql.Open(conf.Db, conf.Conn)	
	defer db.Close()

	res, err := db.Exec("insert into EntryType (Name) values (?)", e.Name)

	if err == nil {
        id, err := res.LastInsertId()
        if err == nil {
            return int(id), nil
        }
    }
	
	return  -1, err
}

func UpdateEntryType(e *models.EntryType) (error) {
	db, err := sql.Open(conf.Db, conf.Conn)	
	defer db.Close()

	_, err = db.Exec("update EntryType set Name = ? where entryTypeId = ?", e.Name, e.ID)

	log.Printf("update " +  e.Name)
	
	return  err
}

func DeleteEntryType(id int) error{
	db, err := sql.Open(conf.Db, conf.Conn)	
	defer db.Close()

	_, err = db.Exec("delete from EntryType where EntryTypeId = ?", id)

	if err == nil {}

	return err
}


func GetEntryTypeById(id int) (*models.EntryType, error) {
	db, err := sql.Open(conf.Db, conf.Conn)	
	defer db.Close()

	rows, err := db.Query("SELECT EntryTypeId, Name FROM entrytype where EntryTypeId = ?", id)

	for rows.Next() {
	    e := new(models.EntryType)
	    if err := rows.Scan(&e.ID, &e.Name); err != nil { }
	    return e, err
	}

	if err != nil {
	}

	return nil, err
}

func ListEntryTypes() ([]*models.EntryType, error) {
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

	return entries, err
}
