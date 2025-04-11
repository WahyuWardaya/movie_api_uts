package controllers

import (
	"movie_api/config"
	"movie_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMovies(c *gin.Context) {
	var movies []models.Movies
	config.DB.Preload("Directors").Preload("Actors").Preload("Genres").Find(&movies)
	c.JSON(http.StatusOK, gin.H{"data": movies})
}

func CreateMovies(c *gin.Context) {
	var movie models.Movies
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan movie utama
	if err := config.DB.Create(&movie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan movie"})
		return
	}

	// Simpan relasi many-to-many
	if len(movie.Actors) > 0 {
		config.DB.Model(&movie).Association("Actors").Replace(movie.Actors)
	}
	if len(movie.Genres) > 0 {
		config.DB.Model(&movie).Association("Genres").Replace(movie.Genres)
	}
	if len(movie.Directors) > 0 {
		config.DB.Model(&movie).Association("Directors").Replace(movie.Directors)
	}

	c.JSON(http.StatusCreated, gin.H{"data": movie})
}

func GetMoviesByID(c *gin.Context) {
	var movie models.Movies
	id := c.Param("id")
	if err := config.DB.Preload("Directors").Preload("Actors").Preload("Genres").First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "movies tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": movie})
}

func UpdateMovies(c *gin.Context) {
	var movie models.Movies
	id := c.Param("id")

	if err := config.DB.Preload("Directors").Preload("Actors").Preload("Genres").First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "movies tidak ditemukan"})
		return
	}

	var updatedMovie models.Movies
	if err := c.ShouldBindJSON(&updatedMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update field utama
	config.DB.Model(&movie).Updates(updatedMovie)

	// Update relasi many-to-many
	config.DB.Model(&movie).Association("Actors").Replace(updatedMovie.Actors)
	config.DB.Model(&movie).Association("Genres").Replace(updatedMovie.Genres)
	config.DB.Model(&movie).Association("Directors").Replace(updatedMovie.Directors)

	c.JSON(http.StatusOK, gin.H{"message": "movies berhasil diperbarui", "data": movie})
}

func DeleteMovies(c *gin.Context) {
	var movie models.Movies
	id := c.Param("id")

	if err := config.DB.Preload("Directors").Preload("Actors").Preload("Genres").First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "movies tidak ditemukan"})
		return
	}

	// Hapus relasi pivot dulu (opsional tapi aman)
	config.DB.Model(&movie).Association("Actors").Clear()
	config.DB.Model(&movie).Association("Genres").Clear()
	config.DB.Model(&movie).Association("Directors").Clear()

	config.DB.Delete(&movie)
	c.JSON(http.StatusOK, gin.H{"message": "movies berhasil dihapus"})
}
