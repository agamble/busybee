package handler

import (
	"errors"
	"net/http"
)

const (
	slackClientID = "315771445571.315870956002"
)

func SlackOAuth(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("error") != "" || r.FormValue("code") == "" {
		failLoudly(w, errors.New("failed to get slack authorization..."))
		return
	}

	_, err := slackConfig().Exchange(r.Context(), r.FormValue("code"))
	if err != nil {
		failLoudly(w, errors.New("failed to exchange slack code"))
		return
	}
}
