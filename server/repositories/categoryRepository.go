package repositories

import (
	"errors"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/lstern/psilibrary/server/models"
	_ "github.com/go-sql-driver/mysql"
)

type CategoryValidator interface{
	GetById(int)(*models.Category, error)
	ValidateCategory(*models.Category) (bool, error)
	CheckForUsedCategory(int) (CategoryCheckResult, error)
	GetByParentId(int)([]*models.Category, error)	
	GetEntriesByCategoryId(int)([]*models.Entry, error)
}

type CategoryCheckResult struct {
	Existing bool
	Categories []*models.Category
	Entries []*models.Entry
}

type CategoryRepository struct{
	Validator CategoryValidator
	Repository
}

func MakeCategoryRepository(validator CategoryValidator, db *sql.DB) CategoryRepository{
	var repo CategoryRepository
	repo.Validator = validator
	repo.DB = db
	return repo
}

func (r CategoryRepository) Create(e *models.Category) (int, error) {
	if r.Validator == nil {
		r.Validator = r
	}

	valid, err := r.Validator.ValidateCategory(e)

	if !valid {
		return -1, err 
	}

	db, err := openSql(r.DB)	
	defer db.Close()

	res, err := db.Exec("insert into category (Name, ParentId) values (?, ?)", e.Name, e.ParentId)

	if err == nil {
        id, err := res.LastInsertId()

        if err != nil {
        	return -1, err
        }

        return int(id), nil
    }
	
	return  -1, err
}

func (r CategoryRepository) Update(e *models.Category) (error) {
	if r.Validator == nil {
		r.Validator = r
	}

	valid, err := r.Validator.ValidateCategory(e)

	if !valid {
		return err 
	}

	db, err := openSql(r.DB)	
	defer db.Close()

	rows, err := db.Exec("update category set Name = ?, ParentId = ? where CategoryId = ?", e.Name, e.ParentId, e.ID)

	if err != nil {
		return err
	}

	count, err := rows.RowsAffected()

	if err == nil && count == 0{
		err = errors.New("Nenhum registro afetado")
	}

	return  err
}

func (r CategoryRepository) Delete(id int) (*CategoryCheckResult, error){
	if r.Validator == nil {
		r.Validator = r
	}

	usedCheck, err := r.Validator.CheckForUsedCategory(id)

	if err != nil {
		return nil, err
	}

	if usedCheck.Existing {
		return &usedCheck, errors.New("Categoria é usada por outros registros")
	}

	db, err := openSql(r.DB)	
	defer db.Close()

	if err != nil {
		return nil, err
	}

	result, err := db.Exec("delete from category where CategoryId = ?", id)

	if err != nil {
		return nil, err
	}

	count, err := result.RowsAffected()

	if err == nil && count == 0{
		err = errors.New("Nenhum registro afetado")
	}

	return nil, err
}


func (r CategoryRepository) GetById(id int) (*models.Category, error) {
	db, err := openSql(r.DB)	
	defer db.Close()

	if err != nil {
		return nil, err
	}

	rows := db.QueryRow("select CategoryId, Name, ParentId FROM category where CategoryId = ?", id)

    e := new(models.Category)
    err = rows.Scan(&e.ID, &e.Name, &e.ParentId)

    return e, err
}

// Verifica se a categoria é valida ou não.
func (r CategoryRepository) ValidateCategory(category *models.Category) (bool, error){
	if r.Validator == nil {
		r.Validator = r
	}

	if category.ParentId != 0{
		cat, err := r.Validator.GetById(category.ParentId)

		if err != nil || cat == nil {
			if cat == nil{
				err = errors.New("Categoria pai não existente")
			}

			return false, err
		}
	}

	return true, nil
}

func (r CategoryRepository) CheckForUsedCategory(id int) (CategoryCheckResult, error){
	var result CategoryCheckResult

	if r.Validator == nil {
		r.Validator = r
	}

	cats, err := r.Validator.GetByParentId(id)

	if err != nil {
		return result, err
	}

	result.Categories = cats

	entries, err := r.Validator.GetEntriesByCategoryId(id)

	if err != nil {
		return result, err
	}

	result.Entries = entries

	result.Existing =  len(cats) > 0 || len(entries) > 0

	return result, err
}

func (r CategoryRepository) List() ([]*models.Category, error) {

	db, err := openSql(r.DB)	
	
	defer db.Close()

	var entries []*models.Category
	rows, err := db.Query("select CategoryId, Name, ParentId FROM category")

	if rows == nil{
		return nil, err
	}

	for rows.Next() {
	    e := new(models.Category)
	    if err := rows.Scan(&e.ID, &e.Name, &e.ParentId); err != nil { }
	    entries = append(entries, e)
	}

	return entries, err
}

// get categories by parentID
func (r CategoryRepository) GetByParentId(catid int)([]*models.Category, error){
	db, err := openSql(r.DB)	
	defer db.Close()

	rows, err := db.Query("select CategoryId, Name, ParentId FROM category where ParentId = ?", catid)

	if rows == nil{
		return nil, err
	}

	var entries []*models.Category
	for rows.Next() {
	    e := new(models.Category)
	    if err := rows.Scan(&e.ID, &e.Name, &e.ParentId); err != nil { }
	    entries = append(entries, e)
	}

	return entries, err
}	

func (r CategoryRepository) GetCategoriesByIdList(ids []int64 ) ([]models.Category, error) {
	db, err := openSql(r.DB)	
	defer db.Close()

	query, args, err := sqlx.In("SELECT * FROM category WHERE CategoryId IN (?);", ids)

	if err != nil {
		return nil, err
	}

	query = db.Rebind(query)

	cats := []models.Category{}	
	err = db.Select(&cats, query, args...)

	return cats, err
}

// get entries from some category
func (CategoryRepository) GetEntriesByCategoryId(catid int)([]*models.Entry, error)	{
	return nil, errors.New("TODO")
}
