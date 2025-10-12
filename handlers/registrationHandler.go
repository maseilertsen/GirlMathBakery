package handlers

import (
	"GirlMathBakery/utils"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

func (s *Server) HandlePostBakery(w http.ResponseWriter, r *http.Request) {
	log.Println("HandlePostBakery Handler")
	// Only handle POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "POST only", http.StatusMethodNotAllowed)
		return
	}

	var br utils.BakeReq

	// Bad json
	if err := json.NewDecoder(r.Body).Decode(&br); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}
	expected := os.Getenv(utils.TOKEN_ENV)
	if expected == "" {
		http.Error(w, "server misconfigured: missing token", http.StatusInternalServerError)
		return
	}
	if br.Token != expected {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	// Bad Token
	// TODO fix token authentication
	if br.Token == "" || br.Token != os.Getenv(utils.TOKEN_ENV) {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	// Ensure quantity
	if br.Item == "" || br.Qty <= 0 {
		http.Error(w, "missing item/qty", http.StatusBadRequest)
		return
	}

	// Handle missing time by time.now()
	when := time.Now()
	if br.Time != "" {
		if t, err := time.Parse(time.RFC3339, br.Time); err == nil {
			when = t
		}
	}

	// Ensure item exists in item table
	_, err := s.DB.Exec(`INSERT OR IGNORE INTO items(name, unit_cost, unit_store, unit) VALUES(?, 0, 0, '')`, br.Item)
	if err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}

	// Insert new bakery into database.
	_, err = s.DB.Exec(`INSERT INTO bakes(when_at, item_name, qty, user)
		                  VALUES(?, ?, ?, ?)`,
		when.UTC().Format(time.RFC3339), br.Item, br.Qty, br.User)
	if err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`ok:true`)) // response message if all is ok
}
