package repositories
/*
import (
	"log"
	"psilibrary/server/models"
	"psilibrary/server/conf"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
*/
/*
func CreateCategory(e *models.Category) (int, error) {
	db, err := sql.Open(conf.Db, conf.Conn)	
	defer db.Close()

	res, err := db.Exec("insert into Category (Name, ParentId) values (?, ?)", e.Name, e.ParentIdD)

	if err == nil {
        id, err := res.LastInsertId()
        if err == nil {
            return int(id), nil
        }
    }
	
	return  -1, err
}

func UpdateCategory(e *models.Category) (error) {
	db, err := sql.Open(conf.Db, conf.Conn)	
	defer db.Close()

	_, err = db.Exec("update Category set Name = ?, ParentId = ? where CategoryId = ?", e.Name, e.ParentId, e.ID)

	log.Printf("update " +  e.Name)
	
	return  err
}

func DeleteCategory(id int) error{
	db, err := sql.Open(conf.Db, conf.Conn)	
	defer db.Close()

	_, err = db.Exec("delete from Category where CategoryId = ?", id)

	if err == nil {}

	return err
}


func GetCategoryById(id int) (*models.Category, error) {
	db, err := sql.Open(conf.Db, conf.Conn)	
	defer db.Close()

	rows, err := db.Query("SELECT CategoryId, Name, ParentId FROM Category where CategoryId = ?", id)

	for rows.Next() {
	    e := new(models.Category)
	    if err := rows.Scan(&e.ID, &e.Name, &e.ParentId); err != nil { }
	    return e, err
	}

	if err != nil {
	}

	return nil, err
}

func ListCategorys() ([]*models.Category, error) {
	var entries []*models.Category

	db, err := sql.Open(conf.Db, conf.Conn)	
	defer db.Close()

	rows, err := db.Query("SELECT CategoryId, Name, ParentId FROM Category")

	for rows.Next() {
	    e := new(models.Category)
	    if err := rows.Scan(&e.ID, &e.Name, &e.ParentId); err != nil { }
	    entries = append(entries, e)
	}

	if err != nil {
	}

	return entries, err
}
*/