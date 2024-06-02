package handlers

import (
	"fmt"
	"net/http"
)

// handlers contains the fiber handlers for the web service.

func ListRecipes(w http.ResponseWriter, r *http.Request) {
	// process request parameters
	// call recipes service
	// return response
	fmt.Fprintln(w, "ListRecipes")
}
