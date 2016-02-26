package repositories

import (
	"log"
	"database/sql"
	"github.com/lstern/psilibrary/server/models"
	_ "github.com/go-sql-driver/mysql"
)

type EntryTypeRepository struct{
	Repository
}

func MakeEntryTypeRepository(db *sql.DB) EntryTypeRepository{
	var repo EntryTypeRepository
	repo.DB = db
	return repo
}

func (r EntryTypeRepository) Create(e *models.EntryType) (int, error) {
	db, err := openSql(r.DB)	
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

func (r EntryTypeRepository) Update(e *models.EntryType) (error) {
	db, err := openSql(r.DB)	
	defer db.Close()

	_, err = db.Exec("update EntryType set Name = ? where entryTypeId = ?", e.Name, e.ID)

	log.Printf("update " +  e.Name)
	
	return  err
}

func (r EntryTypeRepository) Delete(id int) error{
	db, err := openSql(r.DB)	
	defer db.Close()

	_, err = db.Exec("delete from EntryType where EntryTypeId = ?", id)

	if err == nil {}

	return err
}


func (r EntryTypeRepository) GetById(id int) (*models.EntryType, error) {
	db, err := openSql(r.DB)	
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

func (r EntryTypeRepository) List() ([]*models.EntryType, error) {
	var entries []*models.EntryType

	db, err := openSql(r.DB)	
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
