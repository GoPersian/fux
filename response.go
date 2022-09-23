package fux

import "net/http"

type (
	ResponseWriter struct {
		HttpResponseWriter http.ResponseWriter
	}
)

func Response(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		HttpResponseWriter: w,
	}
}

func (w *ResponseWriter) Status(statusCode int) *ResponseWriter {
	w.HttpResponseWriter.WriteHeader(statusCode)
	return w
}
