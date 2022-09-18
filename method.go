package fux

import "net/http"

type (
	Handler func(w http.ResponseWriter, request *http.Request)
)

func (f *Fux) Get(pattern string, handler Handler) {
	f.Router.HandleFunc(pattern, handler).Methods(http.MethodGet)
}

func (f *Fux) Post(pattern string, handler Handler) {
	f.Router.HandleFunc(pattern, handler).Methods(http.MethodPost)
}

func (f *Fux) Put(pattern string, handler Handler) {
	f.Router.HandleFunc(pattern, handler).Methods(http.MethodPut)
}

func (f *Fux) Patch(pattern string, handler Handler) {
	f.Router.HandleFunc(pattern, handler).Methods(http.MethodPatch)
}

func (f *Fux) Delete(pattern string, handler Handler) {
	f.Router.HandleFunc(pattern, handler).Methods(http.MethodDelete)
}
