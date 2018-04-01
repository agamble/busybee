package handler

var debug = false

var ReloadAuth func()

type page struct {
}

func Init(d bool) {
	debug = d
	initTemplates(d)
}
