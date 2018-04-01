package handler

import (
	"net/http"

	"github.com/agamble/page/store"
	"github.com/gorilla/mux"
)

func Heartbeat(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]

	u, err := store.UserByUUID(uuid)
	if err != nil {
		failLoudlyPage(w, err)
		return
	}

	if err := store.WriteHeartbeat(u); err != nil {
		failLoudlyPage(w, err)
		return
	}

	w.Write([]byte("ok"))
}
