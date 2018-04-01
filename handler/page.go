package handler

import (
	"log"
	"net/http"

	"github.com/agamble/page/store"
	"github.com/gorilla/mux"
)

const (
	msgHeader = "X-msg"
)

func Page(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]

	u, err := store.UserByUUID(uuid)
	if err != nil {
		failLoudlyPage(w, err)
		return
	}

	if err := u.CanSendPage(); err != nil {
		failLoudlyPage(w, err)
		return
	}

	msg := r.Header.Get(msgHeader)
	if msg == "" {
		msg = "You are being paged by TESORO"
	}

	w.Write([]byte("ok"))
}

func failLoudlyPage(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("error: failed to send page."))
	log.Printf("Failed to send page: %v", err)
}
