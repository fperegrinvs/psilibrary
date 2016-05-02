package repositories

import (
	"github.com/lstern/psilibrary/server/models"
)

type UserRepository struct{
	Repository
}

func (r UserRepository) GetById(id string) (*models.User, error) {
	db, err := openSql(r.DB)

	if err != nil {
		return nil, err
	}	

	defer db.Close()

	result := models.User{}
	err = db.Get(&result, "SELECT * FROM user where Login = ? LIMIT 1", id)


	if err != nil {
		return nil, err
	}

	return &result, nil;
}
