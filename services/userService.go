package services

import (
	"errors"

	"github.com/RohanSinghCode/gin-practice/config"
	"github.com/RohanSinghCode/gin-practice/internal/database"
	"github.com/RohanSinghCode/gin-practice/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func InsertUser(userRequest models.User, c *gin.Context) (uuid.UUID, error) {
	if userRequest.FirstName == "" {
		return uuid.Nil, errors.New("first Name is required")
	}

	if userRequest.EmailAddress == "" {
		return uuid.Nil, errors.New("email Address is required")
	}

	if userRequest.Password == "" {
		return uuid.Nil, errors.New("password is required")
	}

	connection := config.DbConnection{
		ConnectionString: "postgres://postgres:Newp@ssword8@localhost:5432/GinPractice?sslmode=disable",
	}

	conn, err := connection.ConnectToDb()

	if err != nil {
		return uuid.Nil, errors.New(err.Error())
	}

	defer conn.Close()
	user, err := connection.DB.InsertUser(c, database.InsertUserParams{
		ID:            uuid.New(),
		Firstname:     userRequest.FirstName,
		Lastname:      userRequest.LastName,
		Emailaddresss: userRequest.EmailAddress,
		Password:      userRequest.Password,
	})

	if err != nil {
		return uuid.Nil, errors.New(err.Error())
	}

	return user.ID, nil
}

func GetUser(userId uuid.UUID, c *gin.Context) (models.User, error) {
	connection := config.DbConnection{
		ConnectionString: "postgres://postgres:Newp@ssword8@localhost:5432/GinPractice?sslmode=disable",
	}

	conn, err := connection.ConnectToDb()
	if err != nil {
		return models.User{}, errors.New(err.Error())
	}
	defer conn.Close()

	user, err := connection.DB.GetUserById(c, userId)

	if err != nil {
		return models.User{}, errors.New(err.Error())
	}

	userResponse := models.User{
		Id:           user.ID,
		FirstName:    user.Firstname,
		LastName:     user.Lastname,
		EmailAddress: user.Emailaddresss,
	}

	return userResponse, nil
}
