package routers

import "github.com/gin-gonic/gin"

func Serve() {
	engine := gin.Default()
	router(engine)
}

func router(c *gin.Engine) {

}
