package main

import (
	"GirlMathBakery/utils"
	"database/sql"
	"fmt"
)

func main() {
	fmt.Println("Hello world!") // todo remove
	db, err := sql.Open("sqlite3", utils.dbFile)
}
