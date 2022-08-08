package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dto "github.com/philaden/xm-go-challenge/src/application/dtos"
)

// @Summary Create a new user
// @Description This endpoint creates a new user.
// @Produce json
// @Router /users [post]
func (apiService ApiService) HandleCreateUser(c *gin.Context) {

	var json dto.CreateUserDto

	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Incorrect details supplied, please try again.", "data": nil, "status": false, "error": ""})
		return
	}

	response, message, err := apiService.UserService.CreateUser(json)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": message, "data": response, "status": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User successfully created.",
		"data":    response,
		"status":  true,
		"error":   nil,
	})
}

// @Summary Attempts to get all users
// @Produce json
// @Description This endpoint fetches a list of all users
// @Router /api/users [get]
func (apiService ApiService) GetUsers(c *gin.Context) {

	response, message, err := apiService.UserService.GetUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": message, "data": nil, "status": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Operation Successful",
		"data":    response,
		"status":  true,
		"error":   nil,
	})
}

// @Summary Attempts to get an existing user by email
// @Produce json
// @Description This endpoint fetches a user by email
// @Router /api/users/:email [get]
func (apiService ApiService) HandleGetUserByEmail(c *gin.Context) {

	email, ok := c.GetQuery("email")

	if ok && email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No email found, please try again."})
		return
	}

	response, message, err := apiService.UserService.GetUserByEmail(email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": message, "data": nil, "status": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Operation Successful",
		"data":    response,
		"status":  true,
		"error":   nil,
	})
}
