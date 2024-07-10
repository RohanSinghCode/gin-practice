package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/RohanSinghCode/gin-practice/models"
	"github.com/RohanSinghCode/gin-practice/services"
)

func PostUser(c *gin.Context) {
	var userRequest models.User
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := services.InsertUser(userRequest, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"userId": userId})
}
