package repository

import (
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primary_key"`
	Password  string
	Name      string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

var DB = os.Getenv("POCKAMENT_DB_URL")

func RepositoryMain() {
	db, err := gorm.Open(mysql.Open(DB), &gorm.Config{})
	if err != nil {
		panic("DB Connection Failed")
	}
	db.AutoMigrate(&User{})
}
