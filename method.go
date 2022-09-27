package fux

import (
	"github.com/gorilla/mux"
	"net/http"
)

type (
	Handler func(w http.ResponseWriter, request *http.Request)
)

// Get defines a route with GET method.
func (f *Fux) Get(pattern string, handler Handler) *mux.Route {
	return f.Router.HandleFunc(pattern, handler).Methods(http.MethodGet)
}

// Post defines a route with POST method.
func (f *Fux) Post(pattern string, handler Handler) *mux.Route {
	return f.Router.HandleFunc(pattern, handler).Methods(http.MethodPost)
}

// Put defines a route with PUT method.
func (f *Fux) Put(pattern string, handler Handler) *mux.Route {
	return f.Router.HandleFunc(pattern, handler).Methods(http.MethodPut)
}

// Patch defines a route with PATCH method.
func (f *Fux) Patch(pattern string, handler Handler) *mux.Route {
	return f.Router.HandleFunc(pattern, handler).Methods(http.MethodPatch)
}

// Delete defines a route with DELETE method.
func (f *Fux) Delete(pattern string, handler Handler) *mux.Route {
	return f.Router.HandleFunc(pattern, handler).Methods(http.MethodDelete)
}

// HandleFunc defines a route without method.
// You can define methods like this HandleFunc(...).methods(http.MethodGet,...)
func (f *Fux) HandleFunc(pattern string, handler Handler) *mux.Route {
	return f.Router.HandleFunc(pattern, handler)
}

// Handle registers a new route with a matcher for the URL path.
func (f *Fux) Handle(pattern string, handler http.Handler) *mux.Route {
	return f.Router.Handle(pattern, handler)
}

func (f *Fux) FileServer(pattern string, root http.FileSystem) *mux.Route {
	return f.Handle(pattern, http.FileServer(root))
}
