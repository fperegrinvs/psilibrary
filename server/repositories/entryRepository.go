package repositories

import (
	"errors"
	"github.com/lstern/psilibrary/server/models"
	//"github.com/lstern/psilibrary/server/conf"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type EntryRepository struct{}

type EntryValidator interface{
	ValidateEntry(*models.Entry, EntryValidator) (bool, error, string)
}

func (EntryRepository) CreateEntry(e *models.Entry, mydb *sql.DB, validator EntryValidator) (int, error) {
	return -1, errors.New("TODO")
}

func (EntryRepository) ValidateEntry(e *models.Entry, validator EntryValidator) (bool, error, string) {
	return false, errors.New("TODO"), ""
}

func (EntryRepository) UpdateEntry(e *models.Entry, mydb *sql.DB, validator EntryValidator) (error) {
	return errors.New("TODO")
}

func (EntryRepository) ListEntries(db *sql.DB) ([]*models.Category, error) {
	return nil, errors.New("TODO")
}

func (EntryRepository) GetEntryById(id int, mydb *sql.DB) (*models.Entry, error) {
	return nil, errors.New("TODO")
}

func (EntryRepository) DeleteEntry(id int, mydb *sql.DB) (error){
	return errors.New("TODO")
}

