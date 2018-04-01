package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/agamble/page/store"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/slack"
)

var debug = false

var ReloadAuth func()

type page struct {
}

func Init(d bool) {
	debug = d
	initTemplates(d)
}

type AuthMiddleware struct {
	uuids map[string]bool // [uuid]

	mu sync.RWMutex
}

func (am *AuthMiddleware) Reload() {
	am.mu.Lock()
	defer am.mu.Unlock()
}

func (am *AuthMiddleware) load() {
	us, err := store.AllUsers()
	if err != nil {
		log.Fatal(err)
	}

	for _, u := range us {
		am.uuids[u.UUID] = true
	}
}

func (am *AuthMiddleware) Auth(token string) bool {
	am.mu.RLock()
	defer am.mu.RUnlock()
	return am.uuids[token]
}

func (am *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")
		if am.Auth(token) {
			log.Printf("Authenticated user %s\n", token)
			newR := r.WithContext(context.WithValue(r.Context(), "uuid", token))
			next.ServeHTTP(w, newR)
			return
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

func NewAuthMw() *AuthMiddleware {
	am := &AuthMiddleware{
		uuids: make(map[string]bool),
	}

	am.load()

	return am
}

func RequestUUID(r *http.Request) string {
	return r.Context().Value("uuid").(string)
}

func SlackRedirectURL() string {
	suffix := "oauth/slack"
	if debug {
		return fmt.Sprintf("http://localhost:8080/%v", suffix)
	}

	return fmt.Sprintf("aaaaaa%v", suffix)
}

func slackConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     slackClientID,
		Endpoint:     slack.Endpoint,
		RedirectURL:  SlackRedirectURL(),
		Scopes:       []string{"incoming-webhook"},
		ClientSecret: "6d30dcf6d0858ee3f351d9bc8523d911",
	}
}

func (p page) SlackOAuthLink() string {
	return slackConfig().AuthCodeURL("state", oauth2.AccessTypeOffline)
}
