package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/philaden/xm-go-challenge/src/application/helpers"
)

func LocationAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		var restClient helpers.IRestClient = helpers.RestClient{
			HttpClient: &http.Client{Timeout: time.Duration(2) * time.Second}}

		//callerIp := "104.28.60.46" // This is a test IP address in Cyprus

		callerIp := c.ClientIP()

		headers := getIpApiRequestHeaders()

		url := fmt.Sprintf("https://ipapi.co/%s/country_code/", callerIp)

		response := restClient.Get(url, headers)

		message := "Something went wrong while trying to process that, please try again."
		if !response.IsSuccessStatusCode {

			if *response.Error == nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": message, "data": nil, "status": false, "error": errors.New("Unable retrieve response from web service").Error()})
				return
			}
			if *response.Error == nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": message, "data": nil, "status": false, "error": (*response.Error).Error()})
				return
			}
		}

		if !strings.Contains(response.Result, "CY") {
			message = "You are not permitted to perform this operation based on your location"
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": message, "data": nil, "status": false, "error": nil})
		}

		c.Next()
	}
}

func getIpApiRequestHeaders() map[string]string {
	headers := make(map[string]string)
	headers["User-Agent"] = "ipapi.co/#go-v1.5"
	return headers
}
