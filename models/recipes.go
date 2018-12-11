package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Ingredient struct {
	Id          string "json:'id'"
	Description string "json:'description"
	Deleted     int    "json:'deleted"
}

type Recipe struct {
	Id          string "json:'id'"
	Title       string "json:'title'"
	Description string "json:'description'"
	Ingredients []Ingredient
}

//return all the recipes on the database
func GetAllRecipes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Enter all the recipes")

	page := r.URL.Query().Get("page")
	i, _ := strconv.ParseInt(page, 0, 64)

	result := GetAllRecipesDb(i)
	//test
	fmt.Println("Returning Recipes")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

//return only one recipe
func GetRecipe(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	//fmt.Fprintf(w, "receta número"+id)

	result := GetRecipeDb(id)
	fmt.Println("Getting Recipe #" + id)
	json.NewEncoder(w).Encode(result)
}

func SetRecipe(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Enter to setRecipe")

	var recipe Recipe

	err := json.NewDecoder(r.Body).Decode(&recipe)
	Error(err, "Error Decoding in setRecipe")
	//get the lastinsertid
	idRecipe := SetRecipeDb(recipe)
	recipe.Id = idRecipe
	//return the recipe
	json.NewEncoder(w).Encode(recipe)

}

func SearchRecipe(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Enter to search")
	vars := mux.Vars(r)
	recipe_name := vars["recipe_name"]
	result := SearchRecipeDb(recipe_name)
	json.NewEncoder(w).Encode(result)

}

func DeleteRecipe(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Deleting Recipe #" + vars["id"])
	DeleteRecipeDb(id)

}

func UpdateRecipe(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	fmt.Println("Updating Recipe #" + vars["id"])

	var recipe Recipe

	err := json.NewDecoder(r.Body).Decode(&recipe)
	Error(err, "Error Decoding in update")

	UpdateRecipeDb(recipe)

}
