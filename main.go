package main

import (
	"GirlMathBakery/handlers"
	"GirlMathBakery/utils"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // sqlite3
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello world!") // todo remove

	// Open/Init DB
	db, err := sql.Open("sqlite3", utils.DBFILE)
	utils.Must(err)
	utils.Must(utils.InitSchema(db))
	defer utils.Must(db.Close())

	http.HandleFunc(utils.BAKE, handlers.HandlePostBakery)
	http.HandleFunc(utils.ROOT, handlers.HandleRoot)

	log.Printf("Listening on %s...", utils.ADDR)
	log.Fatal(http.ListenAndServe(utils.ADDR, nil))
}
