package routers

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Serve() {
	engine := gin.Default()
	router(engine)
}

func router(c *gin.Engine) {
	log.Fatalf("Not Inplement")
}
