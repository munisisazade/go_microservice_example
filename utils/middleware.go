package utils

import (
	"log"
	"net/http"
	"runtime"
)

func MiddlewareRequest(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(runtime.Version(), r.URL, r.Method)
		w.Header().Add("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	})
}
