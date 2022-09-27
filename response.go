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

// Response creates an instance of ResponseWriter struct.
func Response(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		HttpResponseWriter: w,
	}
}

// Status sends an HTTP response header with the provided status code.
func (w *ResponseWriter) Status(statusCode int) *ResponseWriter {
	w.HttpResponseWriter.WriteHeader(statusCode)
	return w
}

// ContentTypeJson sets Content-Type: application/json; charset=utf-8.
func (w *ResponseWriter) ContentTypeJson() *ResponseWriter {
	w.HttpResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	return w
}

// ContextTypeHtml sets Content-Type: text/html.
func (w *ResponseWriter) ContextTypeHtml() *ResponseWriter {
	w.HttpResponseWriter.Header().Set("Content-Type", "text/html")
	return w
}

// Header creates an instance of Header struct.
func (w *ResponseWriter) Header() *Header {
	return &Header{
		ResponseWriter: w,
	}
}

// Add adds the key, value pair to the header
func (h *Header) Add(key, value string) *ResponseWriter {
	h.ResponseWriter.HttpResponseWriter.Header().Add(key, value)
	return h.ResponseWriter
}

// Set sets the header entries associated with key to the single element value.
func (h *Header) Set(key, value string) *ResponseWriter {
	h.ResponseWriter.HttpResponseWriter.Header().Set(key, value)
	return h.ResponseWriter
}
