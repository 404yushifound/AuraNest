package controllers

import (
    "AuraNest/config"
    "AuraNest/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetTimeSlots(c *gin.Context) {
    var slots []models.TimeSlot
    config.DB.Find(&slots)

    c.JSON(http.StatusOK, gin.H{
        "data": slots,
    })
}
