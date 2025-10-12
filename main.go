package main

import (
	"GirlMathBakery/handlers"
	"GirlMathBakery/utils"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // sqlite3
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello world!") // todo remove

	// Open/Init DB
	db, err := sql.Open("sqlite3", utils.DBFILE)
	utils.Must(err)
	utils.Must(utils.InitSchema(db))
	defer db.Close()

	srv := handlers.NewServer(db, os.Getenv(utils.TOKEN_ENV))

	http.HandleFunc(utils.BAKE, srv.HandlePostBakery)
	http.HandleFunc(utils.DASHBOARD, srv.HandleDashboard)
	http.HandleFunc(utils.SEED, srv.HandleSeed)
	http.HandleFunc(utils.ROOT, handlers.HandleRoot)

	log.Printf("Listening on %s...", utils.ADDR)
	log.Fatal(http.ListenAndServe(utils.ADDR, nil))
}
