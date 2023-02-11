package Controller

import (
	"BunkyRecipeService/repo"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func getAllRecipe(ctx *gin.Context) {
	fmt.Println("Calling repo.getAllRecipe")
	recipes, err := repo.GetAllRecipe()
	if err != nil {
		log.Fatal("Error error getting Recipes List: ", err.Error())
		ctx.IndentedJSON(http.StatusInternalServerError, nil)
	}

	ctx.IndentedJSON(http.StatusOK, recipes)

}

func RecipeRoutes(route *gin.Engine) {
	recipes := route.Group("/recipe")
	{
		recipes.GET("/getRecipes", getAllRecipe)
	}
}
