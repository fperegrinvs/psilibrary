package repositories

import (
	"errors"
	"github.com/lstern/psilibrary/server/models"
	_ "github.com/go-sql-driver/mysql"
)

type EntryRepository struct{
	Validator EntryValidator
	Repository
}

type EntryValidator interface{
	ValidateEntry(*models.Entry) (bool, string, error)
}

func (r EntryRepository) Create(e *models.Entry) (int, error) {
	if r.Validator == nil{
		r.Validator = r
	}

	//valid, err, msg := validator.ValidateEntry()
	db, err := openSql(r.DB)	
	defer db.Close()

	res, err := db.Exec("insert into Entry (Abstract, Author, Content, EntryTypeID, Journal, PublishData, Title) " +
		"values (?, ?, ?, ?, ?, ?, ?)", e.Abstract, e.Author, e.Content, e.EntryType.ID, e.Journal, e.PublishDate,
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

func (EntryRepository) ValidateEntry(e *models.Entry) (bool, string, error) {
	/*_, err := repository.GetCategoryById(e., nil)

	if (err != nil){
		return false, "Categoria inválida", err
	}
*/
	return true, "", nil
}


func (EntryRepository) Update(e *models.Entry) (error) {
	return errors.New("TODO")
}

func (EntryRepository) List() ([]*models.Category, error) {
	return nil, errors.New("TODO")
}

func (EntryRepository) GetById(id int) (*models.Entry, error) {
	return nil, errors.New("TODO")
}

func (EntryRepository) Delete(id int) (error){
	return errors.New("TODO")
}

