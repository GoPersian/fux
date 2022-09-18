package fux

import "github.com/gorilla/mux"

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
