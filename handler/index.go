package handler

import (
	"errors"
	"net/http"
	"net/url"
	"os"
	"time"
)

type indexPage struct {
	page
}

func Index(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "index.tmpl", indexPage{})
}

func Pay(w http.ResponseWriter, r *http.Request) {
	// Do some payment validation

	email := r.FormValue("email")

	if err := signup(email); err != nil {
		failLoudly(w, err)
		return
	}

	executeTemplate(w, "thanks.tmpl", indexPage{})
}

func signup(email string) error {
	c := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := c.PostForm("https://forum.busybee.careers", url.Values{
		"email": []string{email},
	})
	if err != nil || resp.StatusCode != http.StatusOK {
		return errors.New("failed to signup for forum")
	}

	resp, err = c.PostForm("https://busybee-hq.slack.com/api/users.admin.invite", url.Values{
		"email":      []string{email},
		"token":      []string{os.Getenv("SLACK_TOKEN")},
		"set_active": []string{"true"},
	})
	if err != nil || resp.StatusCode != http.StatusOK {
		return errors.New("failed to signup for forum")
	}

	return nil
}
