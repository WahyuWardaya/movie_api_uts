package controllers

import (
	"movie_api/config"
	"movie_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetActors(c *gin.Context) {
	var actors []models.Actors
	config.DB.Find(&actors)
	c.JSON(http.StatusOK, gin.H{"data": actors})
}

func CreateActors(c *gin.Context) {
	var actors models.Actors
	if err := c.ShouldBindJSON(&actors); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&actors)
	c.JSON(http.StatusCreated, gin.H{"data": actors})
}

func GetActorsByID(c *gin.Context) {
	var actors models.Actors
	id := c.Param("id")
	if err := config.DB.First(&actors, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "actors tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": actors})
}

func UpdateActors(c *gin.Context) {
	var actors models.Actors
	id := c.Param("id")

	if err := config.DB.First(&actors, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "actors tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&actors); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&actors)
	c.JSON(http.StatusOK, gin.H{"message": "actors berhasil diperbarui", "data": actors})
}

func DeleteActors(c *gin.Context) {
	var actors models.Actors
	id := c.Param("id")

	if err := config.DB.First(&actors, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "actors tidak ditemukan"})
		return
	}

	config.DB.Delete(&actors)
	c.JSON(http.StatusOK, gin.H{"message": "actors berhasil dihapus"})
}
