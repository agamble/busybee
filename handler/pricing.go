package handler

import "net/http"

func Pricing(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "pricing.tmpl", nil)
}
