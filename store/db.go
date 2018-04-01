package store

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

func InitDB() {
	var err error
	db, err = sqlx.Open("sqlite3", "./page.db")
	if err != nil {
		panic(err)
	}
}

type Base struct {
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`
}
