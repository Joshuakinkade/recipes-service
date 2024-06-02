package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joshuakinkade/recipes-service/web/handlers"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/recipes", handlers.ListRecipes)
	fmt.Print("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
