package handler

import "net/http"

func Login(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "login.tmpl", nil)
}
