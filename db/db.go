package db

import (
	"database/sql"
	"fmt"
)

var Db *sql.DB

const file string = "main.db"
const create string = `
  CREATE TABLE IF NOT EXISTS urls (
  id INTEGER NOT NULL PRIMARY KEY,
	public_id VARCHAR(255) NOT NULL UNIQUE,
  url VARCHAR(255) NOT NULL
  );`

func Init() {
	db, err := sql.Open("sqlite3", file)

	if err != nil {
		fmt.Println(err)
	}
	if _, err := db.Exec(create); err != nil {
		fmt.Println(err)
	}

	Db = db
}
