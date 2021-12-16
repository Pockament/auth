package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pockament/auth/repository"
	"github.com/pockament/auth/service"
)

type user struct {
	Id   uuid.UUID
	Name string
}

func CreateUserHandler(c *gin.Context) {
	var Option repository.Option
	if err := c.ShouldBindJSON(&Option); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"code":    0x0002,
			"reason":  "",
		})
	}
	User, err := service.CreateUser(
		Option.Name,
		Option.Password,
	)
	user := user{
		Id:   User.ID,
		Name: User.Name,
	}
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"code":    0x0002,
		})
	} else {
		c.JSON(201, gin.H{
			"success": true,
			"code":    0x0001,
			"user":    user,
		})
	}
}

func UpdateUser(c *gin.Context) {
	panic("Not Inplemented!")
}

func DeleteUser(c *gin.Context) {
	panic("Not Inplemented!")
}
