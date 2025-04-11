package controllers

import (
	"movie_api/config"
	"movie_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDirectors(c *gin.Context) {
	var directors []models.Directors
	config.DB.Find(&directors)
	c.JSON(http.StatusOK, gin.H{"data": directors})
}

func CreateDirectors(c *gin.Context) {
	var directors models.Directors
	if err := c.ShouldBindJSON(&directors); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&directors)
	c.JSON(http.StatusCreated, gin.H{"data": directors})
}

func GetDirectorsByID(c *gin.Context) {
	var directors models.Directors
	id := c.Param("id")
	if err := config.DB.First(&directors, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "directors tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": directors})
}

func UpdateDirectors(c *gin.Context) {
	var directors models.Directors
	id := c.Param("id")

	if err := config.DB.First(&directors, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "directors tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&directors); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&directors)
	c.JSON(http.StatusOK, gin.H{"message": "directors berhasil diperbarui", "data": directors})
}

func DeleteDirectors(c *gin.Context) {
	var directors models.Directors
	id := c.Param("id")

	if err := config.DB.First(&directors, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "directors tidak ditemukan"})
		return
	}

	config.DB.Delete(&directors)
	c.JSON(http.StatusOK, gin.H{"message": "directors berhasil dihapus"})
}