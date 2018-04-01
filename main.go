package main

import (
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/agamble/bb/handler"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	debug = true
)

var (
	errNoLayout = errors.New("failure to load layout")
)

func getRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.Index)
	r.HandleFunc("/signup", handler.Pay)

	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	f, err := os.OpenFile("./bb.log", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	w := io.MultiWriter(os.Stdout, f)
	lr := handlers.LoggingHandler(w, r)

	return lr
}

func init() {
	handler.Init(debug)
}

func main() {
	r := getRouter()
	http.ListenAndServe(":8080", r)
}
