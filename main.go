package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"pokemon-api/database"
)

func getAllPokemons(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(database.PokemonDb)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(commonMiddleware)
	myRouter.HandleFunc("/pokemons", getAllPokemons).Methods("GET")

	myRouter.HandleFunc("/pokemons", createNewArticle).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	/*
		{
			"ID": "3",
			"Name": "Bulbasaur",
			"Type": "Plant"
		}
	*/
	db := database.Pokemon{}

	err := json.NewDecoder(r.Body).Decode(&db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.PokemonDb = append(database.PokemonDb, db)

}

func main() {
	fmt.Println("Pokemon Rest API")
	handleRequests()
}
