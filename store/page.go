package store

import "time"

const (
	getPagesForUserQuery = `SELECT * FROM page WHERE `
)

type Page struct {
	UserID    int64     `db:"user_id"`
	Message   string    `db:"msg"`
	CreatedAt time.Time `db:"created_at"`
}
