package handlers

import (
	"log"
	"net/http"

	"NasaImage/views"

	"github.com/gorilla/mux"
)

func Handle() {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1/").Subrouter()
	s.HandleFunc("/", views.Home).Methods("GET")
	s.HandleFunc("/images/", views.ImagesList).Methods("GET")
	s.HandleFunc("/images/{date}", views.Images).Methods("GET")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
