package Controller

import (
	"BunkyRecipeService/utils"
	//Models "BunkyRecipeService/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Context contains data about the http request
func getRecipes(ctx *gin.Context) {
	// here is where we do whatever it is we actually want to do
	// i suppose though this would actually be the domain and it would need to be in another

	fmt.Println("Calling utils.GetRecipes")

	recipes, err := utils.GetRecipe(ctx.Params.ByName("id"))
	if err != nil {
		log.Fatal("Error error getting Recipes: ", err.Error())
	}
	fmt.Printf("Read %s row(s) successfully.\n", recipes[0].RecipeName)

	ctx.IndentedJSON(http.StatusOK, recipes[0])
}

func RecipeRoutes(route *gin.Engine) {
	recipes := route.Group("/recipe")
	{
		recipes.GET("/getRecipes/:id", getRecipes)
	}
}
