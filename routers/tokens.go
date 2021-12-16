package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/pockament/auth/service"
)

func RevokeTokenHandler(c *gin.Context) {
	t := c.GetHeader("authorization")
	if t == "" {
		c.JSON(401, gin.H{
			"success": false,
			"code":    0x0002,
			"reason":  "token not found",
		})
		return
	}

	err := service.RevokeToken(t)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"code":    0x0002,
			"reason":  "Failed to revoke Token",
		})
		return
	} else {
		// 作法的には204を返すべきですが仕様が200なのでそうしています
		c.JSON(200, gin.H{
			"success": true,
			"code":    0x0001,
		})
	}
}
