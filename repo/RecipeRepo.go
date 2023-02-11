package repo

import (
	Models "BunkyRecipeService/models"
	"BunkyRecipeService/service"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var SELECT_INGREDIENTS_BY_ID = "select ring.measurementAmount, ring.measurement, ing.name from bunkyrecipedb.recipe_ingredients as ring inner join bunkyrecipedb.ingredients as ing ON ring.ingredientId = ing.id inner join bunkyrecipedb.recipe_list as rList ON rlist.id = ring.recipeId where rlist.id = %s"
var SELECT_RECIPE_BY_ID = "select id, recipeName,isVegan, timeHours,timeMinutes, timeSeconds from bunkyrecipedb.recipe_list where id = %s"
var SELECT_ALL_RECIPES = "select id, recipeName,isVegan, timeHours,timeMinutes, timeSeconds, imgPath from bunkyrecipedb.recipe_list"

var SELECT_INSTRUCTIONS_BY_ID = "select stepInstruction, stepNum from bunkyrecipedb.instructions where recipeId = %s ORDER BY stepNum ASC"

func GetAllRecipe() (Models.RecipeListResponse, error) {
	// so for this we need to get all the recipes
	// what we will have to do is get the top 50 or so recipes and then filte it down from there.
	//But that is such a far away concern for now.

	var recipeResponse = Models.RecipeListResponse{}
	fmt.Println("Get Recipes")
	db, err := openMySql()
	ctx := context.Background()

	tsql := fmt.Sprintf(SELECT_ALL_RECIPES)
	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return recipeResponse, err
	}
	defer rows.Close()

	var recipeList []Models.RecipeResponse
	for rows.Next() {
		var recipeName, id, imgPath string
		var isVegan, timeHours, timeMinutes, timeSeconds int
		var vegan = false
		if isVegan == 1 {
			vegan = true
		}
		err := rows.Scan(&id, &recipeName, &isVegan, &timeHours, &timeMinutes, &timeSeconds, &imgPath)
		if err != nil {
			return recipeResponse, err
		}

		fmt.Println("Get Ingredients")
		var ing, ingError = GetRecipeIngredientsById(id)
		if ingError != nil {
			return recipeResponse, ingError
		}
		fmt.Println("GetRecipeInstructionsById")
		var instruction, insError = GetRecipeInstructionsById(id)
		if insError != nil {
			return recipeResponse, insError
		}

		fmt.Println("imagePath:" + imgPath)

		image, err := service.GetImageByFilePath(imgPath)
		if err != nil {
		}

		recipeWithImage := Models.RecipeResponse{Id: id, RecipeName: recipeName, IsVegan: vegan, TimeHours: timeHours, TimeMinutes: timeMinutes, TimeSeconds: timeSeconds, Ingredients: ing, Instructions: instruction, FoodPic: image}

		recipeList = append(recipeList, recipeWithImage)
	}
	recipeResponse = Models.RecipeListResponse{RecipeList: recipeList}
	return recipeResponse, nil
}

func GetRecipeIngredientsById(recipeId string) ([]Models.Ingredient, error) {
	db, err := sql.Open("mysql", "root:Chester89!@tcp(127.0.0.1:3306)/bunkyrecipedb")
	ctx := context.Background()

	// Check if database is alive.
	if err != nil {
		return nil, err
	}
	ingredientsTsql := fmt.Sprintf(SELECT_INGREDIENTS_BY_ID, recipeId)
	fmt.Println(ingredientsTsql)
	// Execute query
	ingredientsRows, err := db.QueryContext(ctx, ingredientsTsql)
	if err != nil {
		return nil, err
	}
	fmt.Println("Executed Query Contexts")
	defer ingredientsRows.Close()

	var ingredients []Models.Ingredient

	for ingredientsRows.Next() {
		var name, measurementAmount, measurement string

		err := ingredientsRows.Scan(&measurementAmount, &measurement, &name)
		if err != nil {
			return nil, err
		}

		var ingredient = Models.Ingredient{MeasurementAmount: measurementAmount, Measurement: measurement, Name: name}
		ingredients = append(ingredients, ingredient)
	}
	return ingredients, nil
}

func GetRecipeInstructionsById(recipeId string) ([]Models.Instruction, error) {
	db, err := sql.Open("mysql", "root:Chester89!@tcp(127.0.0.1:3306)/bunkyrecipedb")
	ctx := context.Background()

	// Check if database is alive.
	if err != nil {
		return nil, err
	}
	instructionsTsql := fmt.Sprintf(SELECT_INSTRUCTIONS_BY_ID, recipeId)
	fmt.Println(instructionsTsql)
	// Execute query
	instructionsRows, err := db.QueryContext(ctx, instructionsTsql)
	if err != nil {
		return nil, err
	}
	fmt.Println("Executed Instruction Contexts")
	defer instructionsRows.Close()

	var instructions []Models.Instruction

	for instructionsRows.Next() {
		var stepInstruction string
		var stepNum int

		err := instructionsRows.Scan(&stepInstruction, &stepNum)
		if err != nil {
			return nil, err
		}

		var instruction = Models.Instruction{StepInstruction: stepInstruction, StepNo: stepNum}
		instructions = append(instructions, instruction)
	}
	return instructions, nil
}

func InsertRecipe(recipe Models.Recipe) {

	//We need to make an insert and then use SELECT LAST_INSERT_ID();

	// The ingredients might be a mix of existing and nonexistent ingredients.
	// Too restrictive to not let users enter non existent ingredients.
	// New plan. We save food data locally but when adding new ingredients we reach out to calorie tracker API

	fmt.Println("Get Recipes")
	db, err := openMySql()
	ctx := context.Background()

	tsql := fmt.Sprintf(SELECT_ALL_RECIPES)
	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {

	}
	defer rows.Close()
}

func openMySql() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:Chester89!@tcp(127.0.0.1:3306)/bunkyrecipedb")
	return db, err
}
