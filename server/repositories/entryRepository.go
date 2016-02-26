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
	GetCategoriesByIdList([]int)([]models.Category, error)
	GetEntryTypeById(int) (*models.EntryType, error)
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

func (r EntryRepository) GetCategoriesByIdList(ids []int ) ([]models.Category, error) {
	var catRepo = MakeCategoryRepository(nil, r.DB)
	return catRepo.GetCategoriesByIdList(ids)
}

func (r EntryRepository) GetEntryTypeById(id int) (*models.EntryType, error) {
	return MakeEntryTypeRepository(r.DB).GetById(id)
}

func (r EntryRepository) ValidateEntry(e *models.Entry) (bool, string, error) {
	size := len(e.Categories)

	if size > 0 {
		ids := make([]int, len(e.Categories))
		for i,cat  := range e.Categories {
			ids[i] = cat.ID
		}

		cats, err := r.Validator.GetCategoriesByIdList(ids)

		if err != nil {
			return false, err.Error(), err
		}

		if len(cats) < len(ids) {
			return false, "Categoria inexistente", nil
		}
	}

	entryType, err := r.Validator.GetEntryTypeById(e.EntryType.ID)
	
	if (err != nil){
		return false, err.Error(), err
	}

	if entryType == nil {
		return false, "tipo de registro não encontrado", nil
	}

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

