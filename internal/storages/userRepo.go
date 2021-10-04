package storages

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	Save(u *User)
}

type userRepo struct {
}

func NewUserRepo() userRepo {
	return userRepo{}
}

func (uR userRepo) Save(u *User) *User {
	db := Get()
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		log.Println(err)	
	}
	u.Password = string(hashedPwd)
	db.Create(u)
	return u
}
