package handlers

import (
	"encoding/json"
	"exoplanet-service/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func AddExoplanetHandler(w http.ResponseWriter, r *http.Request) {
	var exoplanet models.Exoplanet
	if err := json.NewDecoder(r.Body).Decode(&exoplanet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := models.AddExoplanet(exoplanet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func ListExoplanetsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("ListExoplanetsHandler called")
	exoplanets := models.ListExoplanets()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exoplanets)
}

func GetExoplanetByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	exoplanet, exists := models.GetExoplanetByID(id)
	if !exists {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exoplanet)
}

func UpdateExoplanetHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var updatedExoplanet models.Exoplanet
	if err := json.NewDecoder(r.Body).Decode(&updatedExoplanet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !models.UpdateExoplanet(id, updatedExoplanet) {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeleteExoplanetHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if !models.DeleteExoplanet(id) {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func FuelEstimationHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	crewCapacity := r.URL.Query().Get("crewCapacity")
	if crewCapacity == "" {
		http.Error(w, "Missing crewCapacity parameter", http.StatusBadRequest)
		return
	}
	// convert crewCapacity to int
	var crewCap int
	_, err := fmt.Sscanf(crewCapacity, "%d", &crewCap)
	if err != nil {
		http.Error(w, "Invalid crewCapacity parameter", http.StatusBadRequest)
		return
	}
	fuelCost, err := models.CalculateFuelCost(id, crewCap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"fuelCost": fuelCost})
}
