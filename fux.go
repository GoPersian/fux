package fux

import (
	"github.com/gorilla/mux"
	"net/http"
)

type (
	Fux struct {
		Router *mux.Router
	}
)

func New() *Fux {
	return &Fux{
		Router: mux.NewRouter(),
	}
}

func (f *Fux) Run(port string) {
	http.ListenAndServe(port, f.Router)
}
