package fux

import (
	"github.com/gorilla/mux"
	"log"
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
	log.Fatalln(http.ListenAndServe(port, f.Router))
}
