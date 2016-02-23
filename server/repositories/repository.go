package repositories
import (
	"database/sql"
	"github.com/lstern/psilibrary/server/conf"
)

type Repository struct{
	DB *sql.DB
}

func openSql(db *sql.DB) (*sql.DB, error){
	if (db != nil){
		return db, nil
	}
	
	db, err := sql.Open(conf.Db, conf.Conn)
	return db, err	
}
