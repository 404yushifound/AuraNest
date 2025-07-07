package controllers

import (
	"AuraNest/config"
	"AuraNest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSongs(c *gin.Context) {
	vibeLabel := c.Query("vibe")
	timeLabel := c.Query("time")

	if vibeLabel == "" || timeLabel == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide both 'vibe' and 'time' query parameters"})
		return
	}

	var song models.Song

	err := config.DB.
		Joins("JOIN vibes ON vibes.id = songs.vibe_id").
		Joins("JOIN time_slots ON time_slots.id = songs.time_slot_id").
		Where("vibes.label = ? AND time_slots.label = ?", vibeLabel, timeLabel).
		Preload("Vibe").
		Preload("TimeSlot").
		First(&song).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No song found for this combination"})
		return
	}

	c.JSON(http.StatusOK, song)
}
