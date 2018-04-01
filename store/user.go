package store

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

const (
	personalPlan = "personal"
	proPlan      = "pro"

	maxPagesPersonal = 100
	maxPagesPro      = 1000

	createUserStmt = `INSERT INTO user (uuid, email) VALUES (?, ?)`

	usersToAlertQuery = `
	SELECT * FROM user
          WHERE
              enabled = 1
          AND
              id NOT IN (
                  SELECT user_id FROM heartbeat WHERE created_at > datetime('now', '-10 minutes')
         )`
)

type User struct {
	ID               int64     `db:"id"`
	GithubJSON       string    `db:"github_json"`
	OAuthSource      string    `db:"oauth_source"`
	StripeCustomerID string    `db:"stripe_customer_id"`
	CreatedAt        time.Time `db:"created_at"`

	SafeUser
}

type SafeUser struct {
	Email     string `db:"email"`
	SiteTitle string `db:"site_title"`
	Plan      string `db:"plan"`
	UUID      string `db:"uuid"`

	UserOptions
}

type UserOptions struct {
	PhoneNumber string `db:"phone_number"`
	ShouldCall  bool   `db:"should_call"`
	ShouldSMS   bool   `db:"should_sms"`
}

func AllUsers() ([]User, error) {
	us := []User{}
	err := db.Select(&us, `SELECT * FROM user`)
	return us, err
}

func UserByUUID(id string) (*User, error) {
	u := User{}
	err := db.Get(&u, `SELECT * FROM user WHERE uuid = ?`, id)
	if err != nil {
		log.Println("failed to get user")
		return nil, err
	}
	return &u, nil
}

func SetUserOptions(uuid string, uo *UserOptions) error {
	_, err := db.Exec(`UPDATE user SET phone_number=:phone_number,
		should_call=:should_call,
		should_sms=:should_sms
		WHERE uuid=:uuid`, User{SafeUser: SafeUser{UserOptions: *uo, UUID: uuid}})
	return err
}

func UsersToAlert() ([]User, error) {
	us := []User{}
	err := db.Select(&us, usersToAlertQuery)
	if err != nil {
		log.Println("failed to get users to alert")
		return nil, err
	}
	return us, nil
}

func (u *User) CanSendPage() error {
	pages, err := u.Pages(time.Hour * 24 * 30)
	if err != nil {
		return err
	}

	if u.Plan == personalPlan && len(pages) >= maxPagesPersonal {
		return errors.New("run out of pages on personal plan")
	}

	if u.Plan == proPlan && len(pages) >= maxPagesPro {
		return errors.New("run out of pages on personal plan")
	}

	return nil
}

func (u *User) Pages(t time.Duration) ([]Page, error) {
	var pages []Page
	threshold := time.Now().Add(-t).Unix()
	err := db.Select(&pages, `SELECT * FROM page WHERE user_id = ? AND created_at > ?`, u.ID, threshold)
	if err != nil {
		return nil, err
	}
	return pages, nil
}

func CreateUser(ctx context.Context, t oauth2.TokenSource) (*User, error) {
	id := uuid.New().String()
	_, err := db.Exec(createUserStmt, id, "")
	if err != nil {
		return nil, err
	}

	log.Println("got here")
	return UserByUUID(id)
}
