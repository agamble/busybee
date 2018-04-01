package store

import (
	"log"
	"time"
)

const (
	slackIntegration = "slack"
	phoneIntegration = "phone"
	smsIntegration   = "sms"
)

type Integration struct {
	UserID      int64     `db:"user_id"`
	Type        string    `db:"type"`
	SlackToken  string    `db:"slack_token"`
	PhoneNumber string    `db:"phone_number"`
	CreatedAt   time.Time `db:"created_at"`
}

func IntegrationsForUser(u *User) ([]Integration, error) {
	var ins []Integration
	err := db.Select(&ins, `SELECT * FROM integration WHERE user_id = ?`, u.ID)
	if err != nil {
		log.Println("failed to get integrations")
		return nil, err
	}
	return ins, nil
}

func (i Integration) Notify(msg string) {
	switch i.Type {
	case slackIntegration:
	case phoneIntegration:
	case smsIntegration:
	}
}
