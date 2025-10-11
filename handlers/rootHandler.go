package handlers

import (
	"GirlMathBakery/utils"
	"fmt"
	"log"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	log.Println("RootHandler entered")
	// Only handle GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Get only", http.StatusMethodNotAllowed)
		return
	}
	_, err := fmt.Fprintln(w, "Hello World") // TODO provide usage here?
	utils.Must(err)                          // Handle errors
}
