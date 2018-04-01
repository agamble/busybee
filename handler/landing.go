package handler

import "net/http"

func Landing(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "landing.tmpl", nil)
}
