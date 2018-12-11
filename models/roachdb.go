package models

import (
	"database/sql"
	"fmt"
	"strconv"
)

var db *sql.DB
var err error

func InitDB(engine string, configuration string) {
	db, err = sql.Open(engine, configuration)
	Error(err, "Error in the connection to the database")
	//db.close()

}

func Error(err error, message string) {
	if err != nil {
		fmt.Println(message)
		panic(err.Error())
	}
}

func GetAllRecipesDb(page int64) []Recipe {

	result, err := db.Query("select idrecipe,title, SUBSTRING (description, 0, 200) from recipe where idrecipe > $1 ORDER BY idrecipe ASC limit 10", page)
	Error(err, "Error gettig all the recipes")
	list_result := make([]Recipe, 0)
	for result.Next() {
		var r Recipe
		// for each row, scan the result into our tag composite object
		err = result.Scan(&r.Id, &r.Title, &r.Description)
		Error(err, "Error scanning the result in the object")
		list_result = append(list_result, r)

	}

	return list_result
}

func GetRecipeDb(id string) []Recipe {

	result, err := db.Query("select idrecipe,title,description from recipe where idrecipe=$1", id)
	Error(err, "Error gettig all the recipes")
	list_result := make([]Recipe, 0)
	for result.Next() {
		var r Recipe
		// for each row, scan the result into our tag composite object
		err = result.Scan(&r.Id, &r.Title, &r.Description)
		Error(err, "Error scanning the result in the object")

		//query for get the ingredients
		ingredients, err := db.Query("select idingredients,description from ingredients where recipe_idrecipe=$1", r.Id)
		Error(err, "Error getting the ingredients")
		list_ingredients := make([]Ingredient, 0)
		var i Ingredient
		for ingredients.Next() {

			ingredients.Scan(&i.Id, &i.Description)
			//add the ingredients to a list
			list_ingredients = append(list_ingredients, i)
		}
		//add the list of ingredients to the recipe struct
		r.Ingredients = list_ingredients

		list_result = append(list_result, r)

	}
	return list_result
}

func SearchRecipeDb(name_recipe string) []Recipe {

	result, err := db.Query("select idrecipe,title,description from recipe where title ilike concat('%',$1,'%') limit 10", name_recipe)
	Error(err, "Error searching recipes")
	list_result := make([]Recipe, 0)
	for result.Next() {
		var r Recipe
		// for each row, scan the result into our tag composite object
		err = result.Scan(&r.Id, &r.Title, &r.Description)
		Error(err, "Error scanning the result in the object")

		//query for get the ingredients
		ingredients, err := db.Query("select idingredients,description from ingredients where recipe_idrecipe=$1", r.Id)
		Error(err, "Error getting the ingredients")
		list_ingredients := make([]Ingredient, 0)
		var i Ingredient
		for ingredients.Next() {

			ingredients.Scan(&i.Id, &i.Description)
			//add the ingredients to a list
			list_ingredients = append(list_ingredients, i)
		}
		//add the list of ingredients to the recipe struct
		r.Ingredients = list_ingredients
		list_result = append(list_result, r)

	}
	return list_result
}

func SetRecipeDb(recipe Recipe) string {

	lastInsertId := 0
	_ = db.QueryRow("INSERT INTO recipe (title, description) VALUES ($1, $2) RETURNING idrecipe", recipe.Title, recipe.Description).Scan(&lastInsertId)

	//get the lastid
	idRecipe := lastInsertId
	fmt.Println(strconv.Itoa(idRecipe), idRecipe)

	//inserting ingredients
	for i := 0; i < len(recipe.Ingredients); i++ {

		_, err := db.Exec("INSERT INTO ingredients (description,recipe_idrecipe) VALUES ($1, $2)", recipe.Ingredients[i].Description, idRecipe)

		if err != nil {
			fmt.Println("error en sql")
			panic(err.Error())
		}

	}
	//db.Close()

	return strconv.Itoa(idRecipe)
}

func DeleteRecipeDb(idRecipe string) bool {

	_, err = db.Exec("Delete from ingredients where recipe_idrecipe=$1", idRecipe)
	Error(err, "Error deleting ingredient")
	_, err = db.Exec("Delete from recipe where idrecipe=$1", idRecipe)
	Error(err, "Error deleting recipe")
	//go show error if not return anything
	return true
}

func UpdateRecipeDb(recipe Recipe) {

	result, err := db.Exec("UPDATE recipe SET title=$1,description=$2 where idrecipe=$3", recipe.Title, recipe.Description, recipe.Id)
	Error(err, "Error updating recipes")

	result.LastInsertId()

	//updating or inserting new ingredients
	for i := 0; i < len(recipe.Ingredients); i++ {

		//search if the ingredient exist
		result := db.QueryRow("SELECT idingredients from ingredients where recipe_idrecipe=$1 and idingredients=$2 limit 1", recipe.Id, recipe.Ingredients[i].Id)
		Error(err, "Error searching a ingredient to after update")

		var id int
		result.Scan(&id)

		//if id is 0 not exist then i create it
		if id == 0 {
			SetIngredient(recipe.Ingredients[i].Description, recipe.Id)

		} else {

			UpdateIngredient(recipe.Ingredients[i].Description, recipe.Ingredients[i].Id)

		}

		//delete the ingredients
		if recipe.Ingredients[i].Deleted == 1 {
			DeleteIngredientDb(recipe.Ingredients[i].Id)
		}

	}

}

func UpdateIngredient(description string, id string) {

	result, err := db.Exec("UPDATE ingredients SET description=$1 where idingredients=$2", description, id)
	Error(err, "Error updating ingredients")
	result.LastInsertId()

}

func SetIngredient(description string, id string) {

	_, err = db.Exec("INSERT INTO ingredients (description,recipe_idrecipe) VALUES ($1, $2)", description, id)
	Error(err, "Error inserting ingredients")

}

func DeleteIngredientDb(idIngredient string) {

	_, err = db.Exec("Delete from ingredients where idingredients=$1", idIngredient)
	Error(err, "Error deleting ingredients")

}
