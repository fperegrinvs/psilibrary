package repositories
import (
	"database/sql"
	"github.com/lstern/psilibrary/server/conf"
	"github.com/jmoiron/sqlx"
)

type Repository struct{
	DB *sql.DB
}

func openSql(db *sql.DB) (*sqlx.DB, error){
	if (db != nil){
		return sqlx.NewDb(db, conf.Db), nil
	}

	db, err := sql.Open(conf.Db, conf.Conn)
	if (err != nil){ 
		return nil, err
	}

	return sqlx.NewDb(db, conf.Db), err	
}
