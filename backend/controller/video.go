package controller

import (
	"net/http"

	"github.com/chanwit/sa-64-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /videos
func CreateVideo(c *gin.Context) {
	var video entity.Video
	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&video).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": video})
}

// GET /video/:id
func GetVideo(c *gin.Context) {
	var video entity.Video

	id := c.Param("id")
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM videos WHERE id = ?", id).Find(&video).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": video})
}

// GET /videos
func ListVideos(c *gin.Context) {
	var videos []entity.Video
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM videos").Find(&videos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": videos})
}

// DELETE /videos/:id
func DeleteVideo(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM videos WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "video not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /videos
func UpdateVideo(c *gin.Context) {
	var video entity.Video
	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", video.ID).First(&video); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "video not found"})
		return
	}

	if err := entity.DB().Save(&video).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": video})
}