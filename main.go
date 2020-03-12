package main

import (
	"NasaImage/db"
	"NasaImage/handlers"
)

func main() {
	defer db.Conn.Close()

	handlers.Handle()
	/*r := mux.NewRouter()*/
	//s := r.PathPrefix("/api/v1/").Subrouter()
	//s.HandleFunc("/", views.Home).Methods("GET")
	//s.HandleFunc("/images/", views.Images).Methods("GET")
	//err = http.ListenAndServe(":8080", r)
	//if err != nil {
	//log.Fatal(err)
	/*}*/
}
