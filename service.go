package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetHandler(w *http.ResponseWriter, r *http.Request) {
}

func PostHandler(w *http.ResponseWriter, r *http.Request) {
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		default:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "{ \"message\": \"not sure we support that buddy\"}")
			return
		case "POST":
			PostHandler(&w, r)
			return
		case "GET":
			GetHandler(&w, r)
			return
		}
	})

	http.ListenAndServe(":3455", r)
}
