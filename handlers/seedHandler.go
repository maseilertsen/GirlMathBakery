package handlers

import (
	"GirlMathBakery/utils"
	"encoding/json"
	"net/http"
	"os"
)

func (s *Server) HandlerSeed(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST only", http.StatusMethodNotAllowed)
		return
	}
	// Auth
	if r.Header.Get("Authorization") != os.Getenv(utils.TOKEN_ENV) {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var payload []struct {
		Item      string  `json:"item"`
		UnitCost  float64 `json:"unit_cost"`
		UnitStore float64 `json:"unit_store"`
		Unit      string  `json:"unit"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}
	tx, err := s.DB.Begin()
	if err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}
	stmt, err := tx.Prepare(`INSERT INTO items(name, unit_cost, unit_store, unit)
		                         VALUES(?,?,?,?)
		                         ON CONFLICT(name) DO UPDATE SET unit_cost=excluded.unit_cost, unit_store=excluded.unit_store, unit=excluded.unit`)
	if err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}

	defer stmt.Close()

	for _, it := range payload {
		if it.Item == "" {
			continue
		}
		if _, err := stmt.Exec(it.Item, it.UnitCost, it.UnitStore, it.Unit); err != nil {
			_ = tx.Rollback()
			http.Error(w, "db error", http.StatusInternalServerError)
			return
		}
	}
	if err := tx.Commit(); err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(`{"ok":true}`))
}
