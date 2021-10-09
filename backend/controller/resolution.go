package controller

import (
	"net/http"

	"github.com/chanwit/sa-64-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /resolutions
func CreateResolution(c *gin.Context) {
	var resolution entity.Resolution
	if err := c.ShouldBindJSON(&resolution); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&resolution).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": resolution})
}

// GET /resolution/:id
func GetResolution(c *gin.Context) {
	var resolution entity.Resolution
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM resolutions WHERE id = ?", id).Scan(&resolution).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resolution})
}

// GET /resolutions
func ListResolutions(c *gin.Context) {
	var resolutions []entity.Resolution
	if err := entity.DB().Raw("SELECT * FROM resolutions").Scan(&resolutions).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resolutions})
}

// DELETE /resolutions/:id
func DeleteResolution(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM resolutions WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "resolution not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /resolutions
func UpdateResolution(c *gin.Context) {
	var resolution entity.Resolution
	if err := c.ShouldBindJSON(&resolution); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", resolution.ID).First(&resolution); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "resolution not found"})
		return
	}

	if err := entity.DB().Save(&resolution).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resolution})
}
