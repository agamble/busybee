package store

import "time"

const (
	createHeartbeatStmt = `INSERT INTO heartbeat (user_id) VALUES (?)`
)

type Heartbeat struct {
	UserID    int64     `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
}

func WriteHeartbeat(u *User) error {
	_, err := db.Exec(createHeartbeatStmt, u.ID)
	return err
}
