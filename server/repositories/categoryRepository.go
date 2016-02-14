package repositories

import (
	"errors"
	//"fmt"
	//"log"
	"psilibrary/server/models"
	"psilibrary/server/conf"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type CategoryValidator interface{
	GetCategoryById(int, *sql.DB)(*models.Category, error)
	ValidateCategory(*models.Category, CategoryValidator) (bool, error)
	CheckForUsedCategory(int, CategoryValidator) (CategoryCheckResult, error)
	GetCategoriesByParentId(int, *sql.DB)([]*models.Category, error)	
	GetEntriesByCategoryId(int, *sql.DB)([]*models.Entry, error)	
}

type CategoryCheckResult struct {
	Existing bool
	Categories []*models.Category
	Entries []*models.Entry
}

type CategoryRepository struct{}

var repository CategoryRepository


func (CategoryRepository) CreateCategory(e *models.Category, mydb *sql.DB, validator CategoryValidator) (int, error) {
	valid, err := validator.ValidateCategory(e, repository)

	if !valid {
		return -1, err 
	}

	db, err := OpenSql(conf.Db, conf.Conn, mydb)	
	defer db.Close()

	res, err := db.Exec("insert into Category (Name, ParentId) values (?, ?)", e.Name, e.ParentId)

	if err == nil {
        id, err := res.LastInsertId()

        if err != nil {
        	return -1, err
        }

        return int(id), nil
    }
	
	return  -1, err
}

func (CategoryRepository) UpdateCategory(e *models.Category, mydb *sql.DB, validator CategoryValidator) (error) {
	valid, err := validator.ValidateCategory(e, repository)

	if !valid {
		return err 
	}

	db, err := OpenSql(conf.Db, conf.Conn, mydb)	
	defer db.Close()

	rows, err := db.Exec("update Category set Name = ?, ParentId = ? where CategoryId = ?", e.Name, e.ParentId, e.ID)

	if err != nil {
		return err
	}

	count, err := rows.RowsAffected()

	if err == nil && count == 0{
		err = errors.New("Nenhum registro afetado")
	}

	return  err
}

func (CategoryRepository) DeleteCategory(id int, mydb *sql.DB, validator CategoryValidator) (*CategoryCheckResult, error){
	usedCheck, err := validator.CheckForUsedCategory(id, validator)

	if err != nil {
		return nil, err
	}

	if usedCheck.Existing {
		return &usedCheck, errors.New("Categoria é usada por outros registros")
	}

	db, err := OpenSql(conf.Db, conf.Conn, mydb)	
	defer db.Close()

	if err != nil {
		return nil, err
	}

	result, err := db.Exec("delete from Category where CategoryId = ?", id)

	if err != nil {
		return nil, err
	}

	count, err := result.RowsAffected()

	if err == nil && count == 0{
		err = errors.New("Nenhum registro afetado")
	}

	return nil, err
}


func (CategoryRepository) GetCategoryById(id int, mydb *sql.DB) (*models.Category, error) {
	db, err := OpenSql(conf.Db, conf.Conn, mydb)	
	defer db.Close()

	rows, err := db.Query("select CategoryId, Name, ParentId FROM Category where CategoryId = ?", id)

	if rows == nil{
		return nil, err
	}

	for rows.Next() {
	    e := new(models.Category)
	    if err := rows.Scan(&e.ID, &e.Name, &e.ParentId); err != nil { }
	    return e, err
	}

	if err != nil {
	}

	return nil, err
}

// Verifica se a categoria é valida ou não.
func (CategoryRepository) ValidateCategory(category *models.Category, getter CategoryValidator) (bool, error){
	if category.ParentId != 0{
		cat, err := getter.GetCategoryById(category.ParentId, nil)

		if err != nil || cat == nil {
			if cat == nil{
				err = errors.New("Categoria pai não existente")
			}

			return false, err
		}
	}

	return true, nil
}

func (CategoryRepository) CheckForUsedCategory(id int, validator CategoryValidator) (CategoryCheckResult, error){
	var result CategoryCheckResult

	cats, err := validator.GetCategoriesByParentId(id, nil)

	if err != nil {
		return result, err
	}

	result.Categories = cats

	entries, err := validator.GetEntriesByCategoryId(id, nil)

	if err != nil {
		return result, err
	}

	result.Entries = entries

	result.Existing =  len(cats) > 0 || len(entries) > 0

	return result, err
}

func (CategoryRepository) ListCategories(db *sql.DB) ([]*models.Category, error) {

	db, err := OpenSql(conf.Db, conf.Conn, db)	
	
	defer db.Close()

	var entries []*models.Category
	rows, err := db.Query("select CategoryId, Name, ParentId FROM Category")

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
func (CategoryRepository) GetCategoriesByParentId(catid int, mydb *sql.DB)([]*models.Category, error){
	db, err := OpenSql(conf.Db, conf.Conn, mydb)	
	defer db.Close()

	rows, err := db.Query("select CategoryId, Name, ParentId FROM Category where ParentId = ?", catid)

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

// get entries from some category
func (CategoryRepository) GetEntriesByCategoryId(catid int, db *sql.DB)([]*models.Entry, error)	{
	return nil, nil
}
