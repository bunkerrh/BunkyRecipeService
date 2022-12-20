package utils

import (
	Models "BunkyRecipeService/models"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var SELECT_INGREDIENTS_BY_ID = "select ring.measurementAmount, ring.measurement, ing.name from bunkyrecipedb.recipe_ingredients as ring inner join bunkyrecipedb.ingredients as ing ON ring.ingredientsId = ing.id inner join bunkyrecipedb.recipe_list as rList ON ring.ingredientsId = ing.id where rlist.id = %s"
var SELECT_RECIPE_BY_ID = "select id, recipeName,isVegan, timeHours,timeMinutes, timeSeconds from bunkyrecipedb.recipe_list where id = %s"
var SELECT_INSTRUCTIONS_BY_ID = "select stepInstruction, stepNum from bunkyrecipedb.instructions where recipeId = %s ORDER BY stepNum ASC"

func GetRecipe(recipeId string) ([]Models.Recipe, error) {

	fmt.Println("Get Recipes")
	db, err := sql.Open("mysql", "root:Chester89!@tcp(127.0.0.1:3306)/bunkyrecipedb")
	ctx := context.Background()

	// Check if database is alive.
	if err != nil {
		return nil, err
	}

	tsql := fmt.Sprintf(SELECT_RECIPE_BY_ID, recipeId)
	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fmt.Println("Get Ingredients")
	var ing, ingError = GetRecipeIngredientsById(recipeId)
	if ingError != nil {
		return nil, ingError
	}

	fmt.Println("GetRecipeInstructionsById")
	var instruction, insError = GetRecipeInstructionsById(recipeId)
	if insError != nil {
		return nil, insError
	}

	var recipe Models.Recipe
	var recipes []Models.Recipe
	for rows.Next() {
		var recipeName string
		var id, isVegan, timeHours, timeMinutes, timeSeconds int
		var vegan = false
		if isVegan == 1 {
			vegan = true
		}
		// Get values from row.
		err := rows.Scan(&id, &recipeName, &isVegan, &timeHours, &timeMinutes, &timeSeconds)
		if err != nil {
			return nil, err
		}

		fmt.Printf("ID: %d, Recipe Name: %s, Vegan: %s, timeHours: %d. timeMinutes: %d, timeSeconds: %d\n",
			id, recipeName, vegan, timeHours, timeMinutes, timeSeconds)

		recipe = Models.Recipe{Id: id, RecipeName: recipeName, IsVegan: vegan, TimeHours: timeHours, TimeMinutes: timeMinutes, TimeSeconds: timeSeconds, Ingredients: ing, Instructions: instruction}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
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
		var stepNum, stepInstruction string

		err := instructionsRows.Scan(&stepInstruction, &stepNum)
		if err != nil {
			return nil, err
		}

		var instruction = Models.Instruction{StepInstruction: stepInstruction, StepNo: stepNum}
		instructions = append(instructions, instruction)
	}
	return instructions, nil
}
