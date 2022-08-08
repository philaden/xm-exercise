package helpers

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claim struct {
	Email  string `json:"email"`
	UserId string `json:"userId"`
	jwt.StandardClaims
}

func CreateSimpleToken(email, userId, secretkey string) (string, error) {
	var signingkey = []byte(secretkey)
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claim{
		Email:  email,
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingkey)

	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func Authenticate(secretkey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "request does not contain an access token"})
			c.Abort()
			return
		}
		err := validateToken(tokenString, secretkey)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}

func validateToken(signedToken, secretkey string) (err error) {
	claims, err := parseToken(signedToken, secretkey)
	if err != nil {
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}

func parseToken(signedToken string, secretkey string) (claims *Claim, err error) {
	signedToken = strings.Split(signedToken, "Bearer ")[1]
	var jwtKey = []byte(secretkey)
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return nil, err
	}
	return claims, nil
}
