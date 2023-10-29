package database

import (
	"url-shortner/repository"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var query = `CREATE TABLE IF NOT EXISTS url (
	id	INTEGER NOT NULL,
	url	TEXT NOT NULL,
	short_url	TEXT NOT NULL UNIQUE,
	PRIMARY KEY("id" AUTOINCREMENT)
)`

var DB *sqlx.DB

func InitializeDB() {
	var err error
	DB, err = sqlx.Open("sqlite3", "url.db")

	if err != nil {
		panic(err)
	}

	_, err = DB.Exec(query)

	if err != nil {
		panic(err)
	}

	repository.UrlRepo = repository.NewUrlRepository(DB)
}
