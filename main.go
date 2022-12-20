package main

import (
	"BunkyRecipeService/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	// Turns on the router and sets the default port to 9090
	router := gin.Default()
	Controller.RecipeRoutes(router)
	router.Run("localhost:9091")
}
