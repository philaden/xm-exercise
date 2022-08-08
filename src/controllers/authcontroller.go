package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dto "github.com/philaden/xm-go-challenge/src/application/dtos"
)

// @Summary logins a user
// @Description This endpoint signs in a user
// @Produce json
// @Router /api/auths [post]
func (apiService ApiService) HandleUserLogin(c *gin.Context) {

	var json dto.LoginDto

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect details supplied, please try again.", "data": nil, "status": false, "error": nil})
		return
	}

	response, message, err := apiService.AuthService.Login(json)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": message, "data": nil, "status": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully logged in",
		"data":    response,
		"status":  true,
		"error":   nil,
	})
}
