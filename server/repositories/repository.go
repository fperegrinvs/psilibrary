package repositories
import (
	"database/sql"
)

type openHandler func(string, string, *sql.DB) (*sql.DB, error)


func OpenSql(dbstr string, conn string, db *sql.DB) (*sql.DB, error){
	if (db != nil){
		return db, nil
	}
	
	db, err := sql.Open(dbstr, conn)
	return db, err	
}
