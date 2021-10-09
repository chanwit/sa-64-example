package controller

import (
	"net/http"
	"time"

	"github.com/chanwit/sa-64-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /watch_videos
func CreateWatchVideo(c *gin.Context) {
	now := time.Now()

	var watchvideo entity.WatchVideo
	var resolution entity.Resolution
	var playlist entity.Playlist
	var video entity.Video

	if err := c.ShouldBindJSON(&watchvideo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", watchvideo.ResolutionID).First(&resolution); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "resolution not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", watchvideo.PlaylistID).First(&playlist); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "playlist not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", watchvideo.VideoID).First(&video); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "video not found"})
		return
	}

	wv := entity.WatchVideo{
		WatchedTime: now,
		Resolution:  resolution,
		Video:       video,
		Playlist:    playlist}

	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wv})
}

// GET /watchvideo/:id
func GetWatchVideo(c *gin.Context) {
	var watchvideo entity.WatchVideo
	id := c.Param("id")
	if err := entity.DB().Preload("Resolution").Preload("Playlist").Preload("Video").Raw("SELECT * FROM watch_videos WHERE id = ?", id).Find(&watchvideo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": watchvideo})
}

// GET /watch_videos
func ListWatchVideos(c *gin.Context) {
	var watchvideos []entity.WatchVideo
	if err := entity.DB().Preload("Resolution").Preload("Playlist").Preload("Video").Raw("SELECT * FROM watch_videos").Find(&watchvideos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": watchvideos})
}

// DELETE /watch_videos/:id
func DeleteWatchVideo(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM watch_videos WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "watchvideo not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /watch_videos
func UpdateWatchVideo(c *gin.Context) {
	var watchvideo entity.WatchVideo
	if err := c.ShouldBindJSON(&watchvideo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", watchvideo.ID).First(&watchvideo); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "watchvideo not found"})
		return
	}

	if err := entity.DB().Save(&watchvideo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": watchvideo})
}
