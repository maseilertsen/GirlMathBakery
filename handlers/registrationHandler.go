package handlers

import "net/http"

func HandlePostBakery(w http.ResponseWriter, r *http.Request) {

	// Only handle POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "POST only", http.StatusMethodNotAllowed)
		return
	}

	// todo Handle bad json
	// todo Handle bad Token
	// todo Handle/ensure quantity/item
	// todo Handle missing time by time.now()

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`ok:true`)) // response message if all is ok
}
