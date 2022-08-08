package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	dto "github.com/philaden/xm-go-challenge/src/application/dtos"
)

// @Summary company registration
// @Produce json
// @Description This endpoint registers a new company
// @Router /api/company [post]
func (apiService ApiService) HandleCompanyRegistration(c *gin.Context) {

	var json dto.CreateCompanyDto

	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Incorrect details supplied, please try again.", "data": nil, "status": false, "error": err.Error()})
		return
	}

	response, message, err := apiService.CompanyService.RegisterCompany(json)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": message, "data": response, "status": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    response,
		"status":  true,
		"error":   nil,
	})
}

// @Summary Attempts to remove an existing company
// @Produce json
// @Description This endpoint deletes an existing company
// @Router /api/company/:id [delete]
func (apiService ApiService) HandleDeleteCompany(c *gin.Context) {

	id := c.Param("id")

	idParam, err := strconv.Atoi(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Something went wrong while trying to process that, please try again.", "data": nil, "status": false, "error": err.Error()})
		return
	}

	response, message, err := apiService.CompanyService.DeleteCompany(idParam)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": message, "data": response, "status": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    response,
		"status":  response,
		"error":   nil,
	})
}

// @Summary Attempts to get all company objects
// @Produce json
// @Description This endpoint fetches a list of all companies
// @Router /api/company [get]
func (apiService ApiService) HandleGetCompanies(c *gin.Context) {

	response, message, err := apiService.CompanyService.GetCompanies(c.Request.URL.Query())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": message, "data": response, "status": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    response,
		"status":  true,
		"error":   nil,
	})
}

// @Summary Retrieves a company object by the company code
// @Produce json
// @Description This endpoint fetches a company object by the company code
// @Router /api/company/:code [get]
func (apiService ApiService) HandleGetCompanyByCompanyCode(c *gin.Context) {

	code := c.Param("code")

	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No code name found, please try again.", "data": nil, "status": false, "error": ""})
		return
	}

	response, message, err := apiService.CompanyService.GetCompanyByCode(code)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": message, "data": response, "status": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Operation Successful",
		"data":    response,
		"status":  true,
		"error":   "",
	})

}
