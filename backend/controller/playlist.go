package controller

import (
	"net/http"

	"github.com/chanwit/sa-64-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /playlists
func CreatePlaylist(c *gin.Context) {
	var playlist entity.Playlist
	if err := c.ShouldBindJSON(&playlist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&playlist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": playlist})
}

// GET /playlist/:id
func GetPlaylist(c *gin.Context) {
	var playlist entity.Playlist
	id := c.Param("id")
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM playlists WHERE id = ?", id).Find(&playlist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": playlist})
}

// GET /playlists
func ListPlaylists(c *gin.Context) {
	var playlists []entity.Playlist
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM playlists").Find(&playlists).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": playlists})
}

// DELETE /playlists/:id
func DeletePlaylist(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM playlists WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "playlist not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /playlists
func UpdatePlaylist(c *gin.Context) {
	var playlist entity.Playlist
	if err := c.ShouldBindJSON(&playlist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", playlist.ID).First(&playlist); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "playlist not found"})
		return
	}

	if err := entity.DB().Save(&playlist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": playlist})
}
