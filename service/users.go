package service

import (
	"errors"
	"github.com/pockament/auth/repository"
)

func CreateUser(Name string, Password string) (repository.User, error) {
	hashed := HashPassword(Password)
	User, status := repository.CreateUser(repository.Option{
		Name:     Name,
		Password: hashed,
	})
	if status != true {
		return repository.User{}, errors.New("error")
	} else {
		return User, nil
	}
}
