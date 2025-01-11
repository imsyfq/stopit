package controllers

import (
	"errors"
	"fmt"
	"stopit/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Credential struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

// var secretKey = []byte(os.Getenv("JWT_SECRET")) // find a way to do this soon
var secretKey = []byte("DontBeADummiesYouGuys!")

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	fmt.Println(string(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("token invalid")
	}

	return nil
}

func Register(c *gin.Context) {

	thePassword := "test"
	hashed, _ := bcrypt.GenerateFromPassword([]byte(thePassword), 8)

	fmt.Println(string(hashed))
}

func Login(c *gin.Context) {
	var cc Credential
	var user models.User

	if err := c.ShouldBind(&cc); err != nil {
		c.JSON(422, gin.H{"message": "Failed to handle request", "success": false})
		return
	}

	// get user by username
	result := models.DB.Where("username = ?", cc.Username).First(&user)
	if result.Error != nil {
		// handle if record not found
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(422, gin.H{"message": "Invalid username or password", "success": false})
			return
		}

		c.JSON(422, gin.H{"message": result.Error})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cc.Password)); err != nil {
		c.JSON(401, gin.H{"message": "Invalid username or password", "success": false})
		return
	}

	token, err := createToken(user.Username)
	if err != nil {
		c.JSON(500, map[string]any{
			"success": false,
			"message": err,
		})
	}

	c.JSON(200, gin.H{
		"success": true,
		"token":   token,
	})
}
