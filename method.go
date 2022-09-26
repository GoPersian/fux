package fux

import (
	"github.com/gorilla/mux"
	"net/http"
)

type (
	Handler func(w http.ResponseWriter, request *http.Request)
)

func (f *Fux) Get(pattern string, handler Handler) *mux.Route {
	return f.Router.HandleFunc(pattern, handler).Methods(http.MethodGet)
}

func (f *Fux) Post(pattern string, handler Handler) *mux.Route {
	return f.Router.HandleFunc(pattern, handler).Methods(http.MethodPost)
}

func (f *Fux) Put(pattern string, handler Handler) *mux.Route {
	return f.Router.HandleFunc(pattern, handler).Methods(http.MethodPut)
}

func (f *Fux) Patch(pattern string, handler Handler) *mux.Route {
	return f.Router.HandleFunc(pattern, handler).Methods(http.MethodPatch)
}

func (f *Fux) Delete(pattern string, handler Handler) *mux.Route {
	return f.Router.HandleFunc(pattern, handler).Methods(http.MethodDelete)
}

func (f *Fux) HandleFunc(pattern string, handler Handler) *mux.Route {
	return f.Router.HandleFunc(pattern, handler)
}
