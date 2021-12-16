package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func RevokeToken(token string) (RevokedTokens, bool) {
	db, err := gorm.Open(mysql.Open(DB), &gorm.Config{})
	if err != nil {
		panic("DB Connection Failed")
	}

	id, _ := uuid.NewUUID()
	revoked := RevokedTokens{ID: id, Token: token}

	res := db.Create(&revoked)
	if res.Error != nil {
		fmt.Printf("Error: %v", res.Error)
		return RevokedTokens{}, false
	}

	return revoked, true
}

func IsRevoked(token string) bool {
	var revokedtokens *RevokedTokens
	db, err := gorm.Open(mysql.Open(DB), &gorm.Config{})
	if err != nil {
		panic("DB Connection Failed")
	}

	db.First(&revokedtokens, token)

	res := db.Create(&revokedtokens)
	if res.Error != nil {
		fmt.Printf("Error: %v", res.Error)
		return true
	}
	
	return false

}
