package routers

import (
	"github.com/gin-gonic/gin"
)

func Serve() {
	engine := gin.Default()
	router(engine)
}

func router(c *gin.Engine) {
	users := c.Group("/users")
	{
		users.POST("/")
		users.PATCH("/")
		users.DELETE("/")
	}

	tokens := c.Group("/tokens")
	{
		tokens.POST("/")
		tokens.DELETE("/")
	}

	verify := c.Group("/")
	{
		verify.GET("/")
	}

}
