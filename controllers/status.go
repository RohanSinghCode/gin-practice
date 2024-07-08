package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/RohanSinghCode/gin-practice/services"
)

func StatusRouter(c *gin.Context) {
	health, err := services.CheckHealth()
	if err == nil {
		c.JSON(200, gin.H{
			"health": health,
		})
		return
	}
	c.JSON(400, gin.H{
		"error": err,
	})
}
