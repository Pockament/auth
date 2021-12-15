package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Option struct {
	Name     string
	Password string
}

func CreateUser(Option Option) (User, bool) {
	db, err := gorm.Open(mysql.Open(DB), &gorm.Config{})
	if err != nil {
		panic("DB Connection Failed")
	}

	id, _ := uuid.NewUUID()
	user := User{ID: id, Name: Option.Name, Password: Option.Password}

	res := db.Create(&user)
	if res.Error != nil {
		fmt.Printf("Error: %v", res.Error)
		return User{}, false
	}

	return user, true
}
