package database

import (
	"fmt"
	"menagment-app-2/src/config"
	model_user "menagment-app-2/src/model/user"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (p *Postgres) CreateUser(req model_user.CreateUserRequest) error {
	userExists, err := p.userExists(req.Email)
	if err != nil {
		return err
	}

	if userExists {
		return fmt.Errorf("user with email %s already exists", req.Email)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	user := model_user.User{
		Email:     req.Email,
		Password:  string(hashedPassword),
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	res := p.db.Create(&user)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (p *Postgres) userExists(email string) (bool, error) {
	var user model_user.User
	res := p.db.Table(config.AppConfig.UsersTable).Where("email = ?", email).First(&user)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, res.Error
	}

	return true, nil

}
