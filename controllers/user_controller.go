package controllers

import (
	"movie_api/config"
	"movie_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.Users
	config.DB.Preload("Roles").Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(c *gin.Context) {
	var users models.Users
	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&users)
	c.JSON(http.StatusCreated, gin.H{"data": users})
}

func GetUserByID(c *gin.Context) {
	var users models.Users
	id := c.Param("id")
	if err := config.DB.Preload("Roles").First(&users, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "users tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func UpdateUsers(c *gin.Context) {
	var users models.Users
	id := c.Param("id")

	if err := config.DB.Preload("Roles").First(&users, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "users tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&users)
	c.JSON(http.StatusOK, gin.H{"message": "users berhasil diperbarui", "data": users})
}

func DeleteUsers(c *gin.Context) {
	var users models.Users
	id := c.Param("id")

	if err := config.DB.Preload("Roles").First(&users, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "users tidak ditemukan"})
		return
	}

	config.DB.Delete(&users)
	c.JSON(http.StatusOK, gin.H{"message": "users berhasil dihapus"})
}
