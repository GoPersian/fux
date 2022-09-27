package fux

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type (
	// Fux is the framework's instance.
	Fux struct {
		Router *mux.Router
	}
)

// New creates an instance of Fux struct.
func New() *Fux {
	return &Fux{
		Router: mux.NewRouter(),
	}
}

// Run calls http.ListenAndServe.
func (f *Fux) Run(port string) {
	log.Fatalln(http.ListenAndServe(port, f.Router))
}

// Vars returns the route variables for the current request, if any.
func Vars(request *http.Request) map[string]string {
	return mux.Vars(request)
}
