package routers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pockament/auth/repository"
	"github.com/pockament/auth/service"
	"io"
)

type user struct {
	Id   uuid.UUID
	Name string
}
type Option struct {
	Name     string
	Password string
}

type update struct {
	Name     *string
	Password *string
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
		*Option.Name,
		*Option.Password,
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

func UpdateUserHander(c *gin.Context) {
	var u update
	body, _ := io.ReadAll(c.Request.Body)
	token := c.Request.Header.Get("authorization")
	if token == "" {
		c.JSON(400, gin.H{
			"success": false,
			"code":    0x0002,
			"reason":  "token not found",
		})
		return
	}

	err := json.Unmarshal(body, &u)
	if err != nil {
		return
	}

	if u.Name != nil && u.Password != nil {
		updateUser, err := service.UpdateUser(token, u.Name, u.Password)
		if err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"code":    0x0002,
				"reason":  "failed to update your information",
			})
		} else {
			c.JSON(201, gin.H{
				"success": true,
				"code":    0x0001,
				"user":    updateUser,
			})
		}
	} else if u.Name == nil {
		updateUser, err := service.UpdateUser(token, nil, u.Password)
		if err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"code":    0x0002,
				"reason":  "failed to update your information.",
			})
		} else {
			c.JSON(201, gin.H{
				"success": true,
				"code":    0x0001,
				"user":    updateUser,
			})
		}
	} else if u.Password == nil {
		updateUser, err := service.UpdateUser(token, u.Name, nil)
		if err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"code":    0x0002,
				"reason":  "failed to update your information..",
			})
		} else {
			c.JSON(201, gin.H{
				"success": true,
				"code":    0x0001,
				"user":    updateUser,
			})
		}
	} else {

	}
}

func DeleteUser(c *gin.Context) {
	panic("Not Inplemented!")
}
