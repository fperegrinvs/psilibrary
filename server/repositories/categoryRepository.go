package repositories

import (
	"errors"
	"log"
	"psilibrary/server/models"
	"psilibrary/server/conf"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type CategoryGetter interface{
	GetCategoryById(int, *sql.DB)(*models.Category, error)
}

type CategoryValidator interface{
	ValidateCategory(*models.Category, CategoryGetter) (bool, error)	
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
        if err == nil {
            return int(id), nil
        }
    }
	
	return  -1, err
}

func (CategoryRepository) UpdateCategory(e *models.Category) (error) {
	db, err := sql.Open(conf.Db, conf.Conn)	
	defer db.Close()

	_, err = db.Exec("update Category set Name = ?, ParentId = ? where CategoryId = ?", e.Name, e.ParentId, e.ID)

	log.Printf("update " +  e.Name)
	
	return  err
}

func (CategoryRepository) DeleteCategory(id int) error{
	db, err := sql.Open(conf.Db, conf.Conn)	
	defer db.Close()

	_, err = db.Exec("delete from Category where CategoryId = ?", id)

	if err == nil {}

	return err
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
func (CategoryRepository) ValidateCategory(category *models.Category, getter CategoryGetter) (bool, error){
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

	if err != nil {
	}

	return entries, err
}
