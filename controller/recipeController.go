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
func getRecipeById(ctx *gin.Context) {
	fmt.Println("Calling utils.getRecipeById")
	recipes, err := utils.GetRecipe(ctx.Params.ByName("id"))
	if err != nil {
		log.Fatal("Error error getting Recipe: ", err.Error())
		ctx.IndentedJSON(http.StatusInternalServerError, nil)
	}
	fmt.Printf("Read %s row(s) successfully.\n", recipes[0].RecipeName)
	ctx.IndentedJSON(http.StatusOK, recipes[0])
}
func getAllRecipe(ctx *gin.Context) {
	fmt.Println("Calling utils.getAllRecipe")
	recipes, err := utils.GetAllRecipe()
	if err != nil {
		log.Fatal("Error error getting Recipes List: ", err.Error())
		ctx.IndentedJSON(http.StatusInternalServerError, nil)
	}

	ctx.IndentedJSON(http.StatusOK, recipes)

}

func RecipeRoutes(route *gin.Engine) {
	recipes := route.Group("/recipe")
	{
		recipes.GET("/getRecipes/:id", getRecipeById)
		recipes.GET("/getRecipes", getAllRecipe)
	}
}
