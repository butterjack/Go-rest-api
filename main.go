package main

import (
	"github.com/butterjack/go-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", controllers.GetOne)
	r.Run()
}
