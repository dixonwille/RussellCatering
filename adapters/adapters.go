package adapters

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type loggingResponseWriter struct {
	status int
	http.ResponseWriter
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.status = code
	lrw.ResponseWriter.WriteHeader(code)
}

//Adapter is used to apply certain attributes to a handler
type Adapter func(http.Handler) http.Handler

//Adapt adds all the adapters to the handler
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

//Logging adds the logging adapter for the handler
func Logging(l *log.Logger) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			myW := &loggingResponseWriter{-1, w}
			before := time.Now().UTC().Unix()
			h.ServeHTTP(myW, r)
			totalTime := time.Now().UTC().Unix() - before
			l.Println(fmt.Sprintf("[%s]\t[%d]\t[%dms]\t%s", r.Method, myW.status, totalTime, r.Host+r.URL.String()))
		})
	}
}

//Header sets headers for response
func Header(key, value string) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add(key, value)
			h.ServeHTTP(w, r)
		})
	}
}
