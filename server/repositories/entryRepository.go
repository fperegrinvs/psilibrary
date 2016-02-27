package repositories

import (
	"errors"
	"github.com/lstern/psilibrary/server/models"
	_ "github.com/go-sql-driver/mysql"
	//"time"
	//"database/sql"
)

type EntryRepository struct{
	Validator EntryValidator
	Repository
}

type EntryValidator interface{
	ValidateEntry(*models.Entry) (error)
}

func MakeEntryRepository(v EntryValidator) EntryRepository {
	var r EntryRepository
	r.Validator = v
	return r
}

func (r EntryRepository) Create(e *models.Entry) (int, error) {
	if r.Validator == nil{
		r.Validator = r
	}

	err := r.Validator.ValidateEntry(e)

	if err != nil {
		return -1, err
	}

	//valid, err, msg := validator.ValidateEntry()
	db, err := openSql(r.DB)	
	defer db.Close()

	//var date = time.Now()
	res, err := db.Exec("insert into Entry (Abstract, Author, Content, EntryTypeID, Journal, PublishDate, Title) " +
		"values (?, ?, ?, ?, ?, ?, ?)", e.Abstract, e.Author, e.Content, e.EntryType.ID, e.Journal,
		 e.PublishDate.Format("2006-01-02"), e.Title)

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

func (r EntryRepository) ValidateEntry(e *models.Entry) (error) {
	if e.Title == "" {
		return errors.New("o título é obrigatório")
	}

	if e.Abstract == "" {
		return errors.New("O campo resumo é obrigatório")
	}

	repo := MakeEntryTypeRepository(nil)
	entry_type, err := repo.GetById(e.EntryType.ID)

	if err != nil {
		return err
	}

	if entry_type == nil || entry_type.ID != e.EntryType.ID {
		return errors.New("Tipo de registro não encontrado")
	}

	catRepo := MakeCategoryRepository(nil, nil)
	for _, cat := range e.Categories {
		rcat, err := catRepo.GetById(cat.ID)
		if err != nil {
			return err
		}

		if rcat.ID != cat.ID {
			return errors.New("Categoria inválida")
		}
	}


	return nil
}

func (r EntryRepository) Update(e *models.Entry) (error) {
	if r.Validator == nil {
		r.Validator = r
	}

	err := r.Validator.ValidateEntry(e)

	if err != nil {
		return err 
	}

	db, err := openSql(r.DB)	
	defer db.Close()

	tx := db.MustBegin()


	rows, err := tx.Exec("update Entry set Abstract = ?, Author = ?,  Content = ?, EntryTypeID = ?, Journal = ?," +
		" PublishDate = ?, Title = ? where EntryID = ?", e.Abstract, e.Author, e.Content, e.EntryType.ID, e.Journal,
		e.PublishDate, e.Title, e.EntryId)
	
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()

	if err!= nil {
		return err
	}

	count, err := rows.RowsAffected()

	if err == nil && count == 0{
		err = errors.New("Nenhum registro afetado")
	}

	return  err
}

func (EntryRepository) List() ([]*models.Category, error) {
	return nil, errors.New("TODO")
}

func (r EntryRepository) GetById(id int) (*models.Entry, error) {
	db, err := openSql(r.DB)

	if err != nil {
		return nil, err
	}	

	defer db.Close()

	result := models.Entry{}
	err = db.Get(&result, "SELECT * FROM Entry where EntryId = ? LIMIT 1", id)

	return &result, err
}

func (EntryRepository) Delete(id int) (error){

	return errors.New("TODO")
}

