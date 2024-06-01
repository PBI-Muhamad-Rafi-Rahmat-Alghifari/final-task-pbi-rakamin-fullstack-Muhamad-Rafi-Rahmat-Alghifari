package controllers

import (
	"FinalTaskAPI/database"
	"FinalTaskAPI/helpers"
	"FinalTaskAPI/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"

	"errors"

	"gorm.io/gorm"
)

func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	User := models.User{}
	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.User = &User

	err := db.Where("id = ?", userID).First(&Photo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = db.Debug().Create(&Photo).Error
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   "Bad Request",
					"message": err.Error(),
				})
				return
			}
			c.JSON(http.StatusCreated, gin.H{
				"data": Photo,
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "User already has a profile photo",
		})
		return
	}
}

func GetPhotos(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	photos := []models.Photo{}
	err := db.Preload("User").Where("id = ?", userID).Find(&photos).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": photos,
	})
}
