package handlers

import (
	"GirlMathBakery/utils"
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {

	// Only handle GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Get only", http.StatusMethodNotAllowed)
		return
	}
	_, err := fmt.Fprintln(w, "Hello World") // TODO provide usage here?
	utils.Must(err)                          // Handle errors
}
