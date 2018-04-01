package handler

import (
	"net/http"

	"github.com/agamble/page/store"
	"github.com/gorilla/mux"
)

type cityPage struct {
	City *store.City
}

func City(w http.ResponseWriter, r *http.Request) {
	city := mux.Vars(r)["city"]
	c, err := store.CityByShortName(city)
	if err != nil {
		failLoudly(w, err)
		return
	}

	executeTemplate(w, "city.tmpl", cityPage{
		City: c,
	})
}
