package main

import (
    "AuraNest/config"
    "AuraNest/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    config.ConnectDatabase()
    config.SeedData()

    r := gin.Default()
    routes.SetupRoutes(r)
    r.Run(":8080")
}
