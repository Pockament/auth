package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Option struct {
	Name     *string
	Password *string
	Id       uuid.UUID
}

func CreateUser(Option Option) (User, bool) {
	db, err := gorm.Open(mysql.Open(DB), &gorm.Config{})
	if err != nil {
		panic("DB Connection Failed")
	}

	id, _ := uuid.NewUUID()
	user := User{ID: id, Name: *Option.Name, Password: *Option.Password}

	res := db.Create(&user)
	if res.Error != nil {
		fmt.Printf("Error: %v", res.Error)
		return User{}, false
	}

	return user, true
}
func UpdateUser(Option Option) (User, bool) {
	db, err := gorm.Open(mysql.Open(DB), &gorm.Config{})
	if err != nil {
		panic("DB Connection Failed")
	}
	var res *gorm.DB
	var user *User
	if Option.Name != nil && Option.Password != nil {
		// どちらもある
		res = db.Model(&user).Where(User{ID: Option.Id}).Updates(User{Name: *Option.Name, Password: *Option.Password})
		if res.Error != nil {
			fmt.Printf("Error: %v", res.Error)
			return User{}, false
		}
		db.First(&user, Option.Id)
		if res.Error != nil {
			return User{}, false
		}
		return *user, true
	} else if Option.Password != nil && Option.Name == nil {
		// Passwordを更新
		res = db.Model(&user).Where(User{ID: Option.Id}).Updates(User{Password: *Option.Password})
		if res.Error != nil {
			fmt.Printf("Error: %v", res.Error)
			return User{}, false
		}
		db.First(&user, Option.Id)
		if res.Error != nil {
			return User{}, false
		}
		return *user, true
	} else if Option.Name != nil && Option.Password == nil {
		res = db.Model(&user).Where(User{ID: Option.Id}).Updates(User{Name: *Option.Name})
		if res.Error != nil {
			fmt.Printf("Error: %v", res.Error)
			return User{}, false
		}
		db.First(&user, Option.Id)
		if res.Error != nil {
			return User{}, false
		}
		return *user, true
	} else {
		return User{}, false
	}
	return User{}, false
}
