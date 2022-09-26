package fux

import "net/http"

type (
	ResponseWriter struct {
		HttpResponseWriter http.ResponseWriter
	}
	Header struct {
		*ResponseWriter
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

func (w *ResponseWriter) ContentTypeJson() *ResponseWriter {
	w.HttpResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	return w
}

func (w *ResponseWriter) Header() *Header {
	return &Header{
		ResponseWriter: w,
	}
}

func (h *Header) Add(key, value string) *ResponseWriter {
	h.ResponseWriter.HttpResponseWriter.Header().Add(key, value)
	return h.ResponseWriter
}

func (h *Header) Set(key, value string) *ResponseWriter {
	h.ResponseWriter.HttpResponseWriter.Header().Set(key, value)
	return h.ResponseWriter
}
