package handler

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

const templateGlob = "./templates/*.tmpl"

var (
	tmpls templateExecutor
)

type templateExecutor interface {
	ExecuteTemplate(wr io.Writer, name string, data interface{}) error
}

type debugTemplateExecutor struct {
	Glob string
}

func (e debugTemplateExecutor) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	t := template.Must(template.ParseGlob(e.Glob))
	return t.ExecuteTemplate(wr, name, data)
}

type releaseTemplateExecutor struct {
	Template *template.Template
}

func (e releaseTemplateExecutor) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	return e.Template.ExecuteTemplate(wr, name, data)
}

func initTemplates(debug bool) {
	if debug {
		tmpls = debugTemplateExecutor{templateGlob}
	} else {
		tmpls = template.Must(template.ParseGlob(templateGlob))
	}
}

func failLoudly(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Failed to render page"))
	log.Printf("Failed loudly with error: %v", err)
}

func executeTemplate(w io.Writer, name string, data interface{}) error {
	return tmpls.ExecuteTemplate(w, name, data)
}
