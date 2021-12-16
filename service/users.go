package service

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/pockament/auth/repository"
)

type Usertype struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	ModifiedPassword bool      `json:"modified_password"`
}

func CreateUser(Name string, Password string) (Usertype, error) {
	hashed := HashPassword(Password)
	User, status := repository.CreateUser(repository.Option{
		Name:     &Name,
		Password: &hashed,
	})
	if status != true {
		return Usertype{}, errors.New("error")
	} else {
		res := Usertype{
			ID:               User.ID,
			Name:             User.Name,
			ModifiedPassword: true,
		}
		return res, nil
	}
}

func UpdateUser(token string, Name *string, Password *string) (Usertype, error) {

	jwtToken, err := CheckJWTToken(token)
	if err != nil {
		return Usertype{}, err
	}
	var claims Claims
	err = json.Unmarshal([]byte(jwtToken), &claims)
	if err != nil {
		return Usertype{}, err
	}
	Id := uuid.MustParse(claims.ID)

	if Password != nil {
		hashed := HashPassword(*Password)
		User, status := repository.UpdateUser(repository.Option{
			Name:     Name,
			Password: &hashed,
			Id:       Id,
		})

		if status != true {
			return Usertype{}, errors.New("error")
		} else {
			res := Usertype{
				ID:               User.ID,
				Name:             User.Name,
				ModifiedPassword: true,
			}
			return res, nil
		}
	} else {
		User, status := repository.UpdateUser(repository.Option{
			Name:     Name,
			Password: nil,
			Id:       Id,
		})
		if status != true {
			return Usertype{}, errors.New("error")
		} else {
			res := Usertype{
				ID:               User.ID,
				Name:             User.Name,
				ModifiedPassword: false,
			}
			return res, nil
		}
	}
}
