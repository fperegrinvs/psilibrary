package repositories

import (
	"errors"
	"github.com/lstern/psilibrary/server/models"
	"github.com/lstern/psilibrary/server/conf"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type EntryRepository struct{
}

type EntryValidator interface{
	ValidateEntry(*models.Entry, EntryValidator) (bool, string, error)
}

func (EntryRepository) Create(e *models.Entry, mydb *sql.DB, validator EntryValidator) (int, error) {
	//valid, err, msg := validator.ValidateEntry()
	db, err := OpenSql(conf.Db, conf.Conn, mydb)	
	defer db.Close()

	res, err := db.Exec("insert into Entry (Abstract, Author, Content, EntryTypeID, Journal, PublishData, Title) " +
		"values (?, ?, ?, ?, ?, ?, ?)", e.Abstract, e.Author, e.Content, e.EntryTypeId, e.Journal, e.PublishDate,
		e.Title)

	if err == nil {
        id, err := res.LastInsertId()

        if err != nil {
        	return -1, err
        }

        return int(id), nil
    }
	
	return  -1, err
}

func (EntryRepository) ValidateEntry(e *models.Entry, validator EntryValidator) (bool, string, error) {
	/*_, err := repository.GetCategoryById(e., nil)

	if (err != nil){
		return false, "Categoria inválida", err
	}
*/
	return true, "", nil
}


func (EntryRepository) Update(e *models.Entry, mydb *sql.DB, validator EntryValidator) (error) {
	return errors.New("TODO")
}

func (EntryRepository) List(db *sql.DB) ([]*models.Category, error) {
	return nil, errors.New("TODO")
}

func (EntryRepository) GetById(id int, mydb *sql.DB) (*models.Entry, error) {
	return nil, errors.New("TODO")
}

func (EntryRepository) Delete(id int, mydb *sql.DB) (error){
	return errors.New("TODO")
}

