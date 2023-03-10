package Models

//alright so what is  in the Recipe
// we need a name, we need the list of ingredients
// we need the steps inolved
// probably need allergen list (actually this will be beyond mvp)

type Recipe struct {
	Id           string        `json:"id"`
	RecipeName   string        `json:"recipeName"`
	ImgPath      string        `json:"imgPath", omitempty`
	IsVegan      bool          `json:"isVegan"`
	TimeHours    int           `json:"timeHours"`
	TimeMinutes  int           `json:"timeMinutes"`
	TimeSeconds  int           `json:"timeSeconds"`
	Ingredients  []Ingredient  `json:"ingredients, omitempty"`
	Instructions []Instruction `json:"instructions, omitempty"`
}

type RecipeResponse struct {
	Id           string        `json:"id"`
	RecipeName   string        `json:"recipeName"`
	FoodPic      string        `json:"foodPic"`
	IsVegan      bool          `json:"isVegan"`
	TimeHours    int           `json:"timeHours"`
	TimeMinutes  int           `json:"timeMinutes"`
	TimeSeconds  int           `json:"timeSeconds"`
	Ingredients  []Ingredient  `json:"ingredients, omitempty"`
	Instructions []Instruction `json:"instructions, omitempty"`
}

type RecipeListResponse struct {
	RecipeList []RecipeResponse `json:"recipeList"`
}

type Ingredient struct {
	Name              string `json:"name""`
	MeasurementAmount string `json:"measurementAmount"`
	Measurement       string `json:"measurement"`
	Id                int    `json:"id,omitempty"`
	RecipeId          int    `json:"recipeId,omitempty"`
}

type Instruction struct {
	StepNo          int    `json:"stepNo"`
	StepInstruction string `json:"stepInstructions"`
	RecipeId        int    `json:"recipeId,omitempty"`
}
