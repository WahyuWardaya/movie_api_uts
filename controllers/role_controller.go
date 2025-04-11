package controllers

import (
	"movie_api/config"
	"movie_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRoles(c *gin.Context) {
	var roles []models.Roles
	config.DB.Find(&roles)
	c.JSON(http.StatusOK, gin.H{"data": roles})
}

func CreateRoles(c *gin.Context) {
	var roles models.Roles
	if err := c.ShouldBindJSON(&roles); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&roles)
	c.JSON(http.StatusCreated, gin.H{"data": roles})
}

func GetRolesByID(c *gin.Context) {
	var roles models.Roles
	id := c.Param("id")
	if err := config.DB.First(&roles, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "roles tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": roles})
}

func UpdateRoles(c *gin.Context) {
	var roles models.Roles
	id := c.Param("id")

	if err := config.DB.First(&roles, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "roles tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&roles); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&roles)
	c.JSON(http.StatusOK, gin.H{"message": "roles berhasil diperbarui", "data": roles})
}

func DeleteRoles(c *gin.Context) {
	var roles models.Roles
	id := c.Param("id")

	if err := config.DB.First(&roles, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "roles tidak ditemukan"})
		return
	}

	config.DB.Delete(&roles)
	c.JSON(http.StatusOK, gin.H{"message": "roles berhasil dihapus"})
}
