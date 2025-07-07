package controllers

import (
    "AuraNest/config"
    "AuraNest/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetVibes(c *gin.Context) {
    var vibes []models.Vibe
    config.DB.Find(&vibes)

    c.JSON(http.StatusOK, gin.H{
        "data": vibes,
    })
}
