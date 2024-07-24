package main

import (
	"exoplanet-service/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Test route working!")
	}).Methods("GET")
	r.HandleFunc("/exoplanets", handlers.AddExoplanetHandler).Methods("POST")
	r.HandleFunc("/exoplanets", handlers.ListExoplanetsHandler).Methods("GET")
	r.HandleFunc("/exoplanets/{id}", handlers.GetExoplanetByIDHandler).Methods("GET")
	r.HandleFunc("/exoplanets/{id}", handlers.UpdateExoplanetHandler).Methods("PUT")
	r.HandleFunc("/exoplanets/{id}", handlers.DeleteExoplanetHandler).Methods("DELETE")
	r.HandleFunc("/fuel/{id}", handlers.FuelEstimationHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
