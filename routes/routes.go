package routes

import (
    "github.com/gin-gonic/gin"
    "AuraNest/controllers"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/vibes", controllers.GetVibes)
    r.GET("/timeslots", controllers.GetTimeSlots)
    r.GET("/songs", controllers.GetSongs)
}
