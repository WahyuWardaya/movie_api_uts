package controllers

import (
	"movie_api/config"
	"movie_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGenres(c *gin.Context) {
	var genres []models.Genres
	config.DB.Find(&genres)
	c.JSON(http.StatusOK, gin.H{"data": genres})
}

func CreateGenres(c *gin.Context) {
	var genres models.Genres
	if err := c.ShouldBindJSON(&genres); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&genres)
	c.JSON(http.StatusCreated, gin.H{"data": genres})
}

func GetGenresByID(c *gin.Context) {
	var genres models.Genres
	id := c.Param("id")
	if err := config.DB.First(&genres, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "genres tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": genres})
}

func UpdateGenres(c *gin.Context) {
	var genres models.Genres
	id := c.Param("id")

	if err := config.DB.First(&genres, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "genres tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&genres); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&genres)
	c.JSON(http.StatusOK, gin.H{"message": "genres berhasil diperbarui", "data": genres})
}

func DeleteGenres(c *gin.Context) {
	var genres models.Genres
	id := c.Param("id")

	if err := config.DB.First(&genres, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "genres tidak ditemukan"})
		return
	}

	config.DB.Delete(&genres)
	c.JSON(http.StatusOK, gin.H{"message": "genres berhasil dihapus"})
}
