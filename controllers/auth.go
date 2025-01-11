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

type MyClaims struct {
	jwt.StandardClaims
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	Exp      int64  `json:"exp"`
}

// var secretKey = []byte(os.Getenv("JWT_SECRET")) // find a way to do this soon
var secretKey = []byte("DontBeADummiesYouGuys!")

func createToken(username string, user_id int) (string, error) {
	claims := MyClaims{
		UserId:   user_id,
		Username: username,
		Exp:      time.Now().Add(time.Hour * 24).Unix(),
		// Exp: time.Now().Add(time.Second * 30).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
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

	token, err := createToken(user.Username, user.Id)
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
