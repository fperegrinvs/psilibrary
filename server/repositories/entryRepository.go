package repositories

import (
	"errors"
	"github.com/lstern/psilibrary/server/models"
	_ "github.com/go-sql-driver/mysql"
	//"time"
	"github.com/jmoiron/sqlx"
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

func (r EntryRepository) insertEntryCategory(entryId int, categoryId int, transaction *sqlx.Tx) error{
	res, err := transaction.Exec("insert into CategoryEntry (EntryId, CategoryId) values (?, ?)", entryId, categoryId)

	if err == nil {
        id, err := res.RowsAffected()

        if err != nil {
        	return err
        }

        if id != 1 {
        	return errors.New("Categoria não foi inserida")
        }
    }
	
	return  err
}

func (r EntryRepository) InsertEntryCategory(entryId int, categoryId int,) (error) {
	db, err := openSql(r.DB)	
	defer db.Close()

	tx := db.MustBegin()
	r.insertEntryCategory(entryId, categoryId, tx)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (r EntryRepository) Create(e *models.Entry) (int, error) {
	if r.Validator == nil{
		r.Validator = r
	}

	err := r.Validator.ValidateEntry(e)

	if err != nil {
		return -1, err
	}

	db, err := openSql(r.DB)	
	defer db.Close()

	tx := db.MustBegin()

	res, err := tx.Exec("insert into Entry (Abstract, Author, Content, EntryTypeID, Journal, PublishDate, Title, MedlineId) " +
		"values (?, ?, ?, ?, ?, ?, ?, ?)", e.Abstract, e.Author, e.Content, e.EntryType.ID, e.Journal,
		 e.PublishDate.Format("2006-01-02"), e.Title, e.MedlineId)

	if err == nil {
        id, err := res.LastInsertId()

        if err == nil {
        	if len(e.Categories) > 0 {
        		for _, cat := range e.Categories {
        			err = r.insertEntryCategory(int(id), cat.ID, tx)

        			if err != nil {
        				tx.Rollback()
        				return -1, err
        			}
        		}
        	}

	        tx.Commit()
	        return int(id), nil
        }
    }
	
	tx.Rollback()
	return  -1, err
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
	catsMap := map[int]*models.Category{}
	for _, cat := range e.Categories {
		rcat, err := catRepo.GetById(cat.ID)
		if err != nil {
			return err
		}

		if rcat.ID != cat.ID {
			return errors.New("Categoria inválida")
		}

		catsMap[rcat.ID] = &cat
	}

	if len(catsMap) < len(e.Categories) {
		return errors.New("Categoria duplicada")
	}

	if (e.MedlineId != "") {
		o, err := r.GetByMedlineId(e.MedlineId);

		if (err == nil && o.MedlineId == e.MedlineId) {
			return errors.New("MedlineId já existente: " + e.MedlineId);
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
		" PublishDate = ?, Title = ?, MedlineId = ? where EntryID = ?", e.Abstract, e.Author, e.Content, e.EntryType.ID, e.Journal,
		e.PublishDate, e.Title, e.MedlineId, e.EntryId)
	
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.MustExec("delete from categoryEntry where EntryID = ?", e.EntryId)

   	if len(e.Categories) > 0 {
		for _, cat := range e.Categories {
			err = r.insertEntryCategory(e.EntryId, cat.ID, tx)

			if err != nil {
				tx.Rollback()
				return err
			}
		}
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

func (r EntryRepository) List() (*[]models.Entry, error) {
	db, err := openSql(r.DB)

	if err != nil {
		return nil, err
	}	

	defer db.Close()

	result := []models.Entry{}
	err = db.Select(&result, "SELECT * FROM Entry")

	if err != nil {
		return nil, err
	}

	for _, item := range result {
		_, err := r.completeEntry(&item)

		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (r EntryRepository) completeEntry(result *models.Entry) (*models.Entry, error){
	if result.EntryTypeId > 0 {
		entry_type, err := MakeEntryTypeRepository(nil).GetById(result.EntryTypeId)

		if err != nil {
			return nil, err
		}

		result.EntryType = *entry_type
	}

	cats, err := r.GetEntryCategories(result.EntryId)

	if err != nil {
		return nil, err
	}

	if len(cats) > 0 {
		catRepo := MakeCategoryRepository(nil, nil)
		result.Categories, err = catRepo.GetCategoriesByIdList(cats)
	}

	return result, err	
}

func (r EntryRepository) GetById(id int) (*models.Entry, error) {
	db, err := openSql(r.DB)

	if err != nil {
		return nil, err
	}	

	defer db.Close()

	result := models.Entry{}
	err = db.Get(&result, "SELECT * FROM Entry where EntryId = ? LIMIT 1", id)


	if err != nil {
		return nil, err
	}

	return r.completeEntry(&result)
}

func (r EntryRepository) GetByMedlineId(id string) (*models.Entry, error) {
	db, err := openSql(r.DB)

	if err != nil {
		return nil, err
	}	

	defer db.Close()

	result := models.Entry{}
	err = db.Get(&result, "SELECT * FROM Entry where MedlineId = ? LIMIT 1", id)


	if err != nil {
		return nil, err
	}

	return r.completeEntry(&result)
}

func (r EntryRepository) GetEntryCategories(id int) ([]int64, error) {
	db, err := openSql(r.DB)

	if err != nil {
		return nil, err
	}	

	defer db.Close()

	var result []int64
	err = db.Select(&result, "SELECT CategoryId FROM CategoryEntry where EntryId = ?", id)

	return result, err
}

func (r EntryRepository) Delete(id int) (error){
	db, err := openSql(r.DB)

	defer db.Close()

	result, err := db.Exec("Delete from Entry where EntryId = ?", id)
	rows, _ := result.RowsAffected()

	if rows != 1 {
		return errors.New("Nenhum registro foi removido")
	}

	return err
}

