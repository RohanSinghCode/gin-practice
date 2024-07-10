package main

import (
	"github.com/RohanSinghCode/gin-practice/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", controllers.StatusRouter)

	// Users endpoint
	r.POST("/users", controllers.PostUser)
	r.Run()
}
