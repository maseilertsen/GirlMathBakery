package handlers

import (
	"GirlMathBakery/utils"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

func (s *Server) HandleDashboard(w http.ResponseWriter, r *http.Request) {
	log.Println("Entered HandleDashboard")

	type row struct {
		Item    string
		Qty     int
		Make    float64
		Store   float64
		Savings float64
	}
	type recent struct {
		When string
		User string
		Item string
		Qty  int
	}
	var totals struct{ Make, Store, Savings float64 }

	// Compute totals
	utils.Must(s.DB.QueryRow(`
			SELECT
			  IFNULL(SUM(b.qty * i.unit_cost), 0),
			  IFNULL(SUM(b.qty * i.unit_store), 0)
			FROM bakes b
			LEFT JOIN items i ON b.item_name = i.name
		`).Scan(&totals.Make, &totals.Store))
	totals.Savings = totals.Store - totals.Make

	// Per item
	rows, err := s.DB.Query(`
			SELECT b.item_name,
			       SUM(b.qty) as qty,
			       IFNULL(SUM(b.qty * i.unit_cost),0) as make,
			       IFNULL(SUM(b.qty * i.unit_store),0) as store
			FROM bakes b
			LEFT JOIN items i ON b.item_name = i.name
			GROUP BY b.item_name
			ORDER BY (store - make) DESC
		`)
	utils.Must(err)
	defer rows.Close()
	var per []row
	for rows.Next() {
		var r row
		utils.Must(rows.Scan(&r.Item, &r.Qty, &r.Make, &r.Store))
		r.Savings = r.Store - r.Make
		per = append(per, r)
	}

	// Recent (last 20)
	rrows, err := s.DB.Query(`
			SELECT when_at, user, item_name, qty
			FROM bakes ORDER BY when_at DESC LIMIT 20
		`)
	utils.Must(err)
	defer rrows.Close()
	var last []recent
	for rrows.Next() {
		var rr recent
		var when string
		utils.Must(rrows.Scan(&when, &rr.User, &rr.Item, &rr.Qty))
		// Show local-ish short time
		if t, err := time.Parse(time.RFC3339, when); err == nil {
			rr.When = t.Format("2006-01-02 15:04")
		} else {
			rr.When = when
		}
		last = append(last, rr)
	}

	data := map[string]any{
		"Totals":  totals,
		"PerItem": per,
		"Recent":  last,
	}
	tplPath := filepath.Join("templates", "dashboard.html")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		http.Error(w, "template error", http.StatusInternalServerError)
		log.Println("template error:", err)
		return
	}
	utils.Must(tpl.Execute(w, data))
}
