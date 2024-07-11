package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

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
		return
	}

	c.JSON(http.StatusOK, gin.H{"userId": userId})
}

type userIdRequest struct {
	ID string `uri:"userId" binding:"required,uuid"`
}

func GetUser(c *gin.Context) {
	var userRequest userIdRequest
	if err := c.ShouldBindUri(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := uuid.Parse(userRequest.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := services.GetUser(userId, c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
