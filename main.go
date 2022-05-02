package main

import (
	"example/router"

	"github.com/gin-gonic/gin"
)

func main() {
	defaultRouter := gin.Default()
	defaultRouter.GET("/ping", router.GetPong)
	defaultRouter.POST("/products/:id", router.Purchase)
	defaultRouter.Run(":9000") // listen and serve on 0.0.0.0:9000
}
