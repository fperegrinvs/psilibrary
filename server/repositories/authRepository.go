package repositories

import (
	//"log"
	"time"
	"database/sql"
	"github.com/lstern/psilibrary/server/models"
	//_ "github.com/go-sql-driver/mysql"
	"github.com/twinj/uuid"
)

type AuthRepository struct{
	Repository
}

func MakeAuthRepository(db *sql.DB) AuthRepository{
	var repo AuthRepository
	repo.DB = db
	return repo
}

func (a AuthRepository) CreateSession(email string) (*models.Session){
	guid := uuid.NewV4()
	now := time.Now()

	session := new(models.Session)
	session.Login = email
	session.ID = guid.String()
	session.Expiration = now.AddDate(0, 0, 5)
	return session
}

func (a AuthRepository) SaveSession(session *models.Session) error {
	db, err := openSql(a.DB)	
	defer db.Close()

	_, err = db.Exec("insert into Session (sessionId, login, expiration) values (?, ?, ?)", session.ID, session.Login, session.Expiration)
	
	return  err
}

func (a AuthRepository) GetToken(sessionId string) (*models.Token, error) {
	db, err := openSql(a.DB)

	if err != nil {
		return nil, err
	}	

	defer db.Close()

	result := new(models.Token)
	err = db.Get(result, "select user.*, session.SessionId, session.Expiration from session inner join user on session.login = user.login where session.SessionId = ? order by session.Expiration desc limit 1", sessionId)


	if err != nil {
		return nil, err
	}

	return result, nil
}
