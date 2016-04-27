package repositories

import (
	"time"
	"database/sql"
	"net/url"
	"github.com/lstern/psilibrary/server/models"
	"github.com/twinj/uuid"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/objx"
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

func (a AuthRepository) CheckSession(sessionId string, login string, token *models.Token) bool{
	if (token == nil || login != token.Login || sessionId != token.ID || token.Expiration.IsZero())  {
		return false
	}

	if token.Expiration.Before(time.Now()){
		return false
	}

	return true
}

func (a AuthRepository) GetFacebookUrl() (string, error){
	provider, err := gomniauth.Provider("facebook")

	if err != nil {
		return "", err
	}

	state := gomniauth.NewState("after", "success")
	authUrl, err := provider.GetBeginAuthURL(state, nil)

	return authUrl, err

}

func (a AuthRepository) ProcessFacebookCallback(rawurl string) (*common.Credentials, error){
	provider, err := gomniauth.Provider("facebook")
	if err != nil {
		return nil, err
	}

	url_obj, _ := url.Parse(rawurl)


	omap, err := objx.FromURLQuery(url_obj.RawQuery)
	if err != nil {
		return nil, err
	}

	creds, err := provider.CompleteAuth(omap)

	return creds, err
}
