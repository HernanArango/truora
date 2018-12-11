package main

import (
	"fmt"
	"log"
	"net/http"
	"truora/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("server runing!")
	//configuration mysql
	//configuration := "root:@tcp(127.0.0.1:3306)/recipes"
	//engine := "mysql"
	//configuration roach
	configuration := "host=172.17.0.1 port=26257 user=roach dbname=recipes sslmode=disable"
	engine := "postgres"
	models.InitDB(engine, configuration)
	r := mux.NewRouter()
	r.HandleFunc("/", models.GetAllRecipes).Methods("GET", "OPTIONS")
	r.HandleFunc("/recipes", models.GetAllRecipes).Methods("GET", "OPTIONS")
	r.HandleFunc("/recipes/{id}", models.GetRecipe).Methods("GET", "OPTIONS")
	r.HandleFunc("/recipes", models.SetRecipe).Methods("POST", "OPTIONS")
	r.HandleFunc("/recipes/delete/{id}", models.DeleteRecipe).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/recipes/update/{id}", models.UpdateRecipe).Methods("PUT", "OPTIONS")
	r.HandleFunc("/recipes/search/{recipe_name}", models.SearchRecipe).Methods("GET", "OPTIONS")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}), handlers.AllowedOrigins([]string{"*"}))(r)))

}
