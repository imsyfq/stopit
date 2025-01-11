package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"stopit/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

// var secretKey = []byte(os.Getenv("JWT_SECRET")) // find a way to do this soon
var secretKey = []byte("DontBeADummiesYouGuys!")

var User models.User

type MyClaims struct {
	jwt.StandardClaims
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	Exp      int64  `json:"exp"`
}

func validateToken(tokenString string) (*MyClaims, error) {
	var claims MyClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Token invalid")
	}

	return &claims, nil
}

func setLoggedUser(claims *MyClaims) {
	var u models.User
	result := models.DB.Where("id = ?", claims.UserId).First(&u)
	if result.Error != nil {
		// handle if record not found
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			panic("user not found")
		}
		panic(result.Error)
	}

	User = u
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		claims, err := validateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if time.Now().Unix() > claims.Exp {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired, please re-login"})
			c.Abort()
			return
		}

		setLoggedUser(claims)

		c.Next()
	}
}
