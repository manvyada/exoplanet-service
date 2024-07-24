package models

import "errors"

type ExoplanetType string

const (
	GasGiant    ExoplanetType = "GasGiant"
	Terrestrial ExoplanetType = "Terrestrial"
)

type Exoplanet struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Distance    float64       `json:"distance"`
	Radius      float64       `json:"radius"`
	Mass        *float64      `json:"mass,omitempty"`
	Type        ExoplanetType `json:"type"`
}

var exoplanets = make(map[string]Exoplanet)

func AddExoplanet(exoplanet Exoplanet) (string, error) {
	if exoplanet.ID == "" || exoplanet.Name == "" || exoplanet.Description == "" || exoplanet.Distance <= 0 || exoplanet.Radius <= 0 {
		return "", errors.New("missing or invalid fields")
	}
	exoplanets[exoplanet.ID] = exoplanet
	return exoplanet.ID, nil
}

func ListExoplanets() []Exoplanet {
	var result []Exoplanet
	for _, exoplanet := range exoplanets {
		result = append(result, exoplanet)
	}
	return result
}

func GetExoplanetByID(id string) (Exoplanet, bool) {
	exoplanet, exists := exoplanets[id]
	return exoplanet, exists
}

func UpdateExoplanet(id string, updatedExoplanet Exoplanet) bool {
	_, exists := exoplanets[id]
	if !exists {
		return false
	}
	exoplanets[id] = updatedExoplanet
	return true
}

func DeleteExoplanet(id string) bool {
	_, exists := exoplanets[id]
	if !exists {
		return false
	}
	delete(exoplanets, id)
	return true
}

func CalculateFuelCost(id string, crewCapacity int) (float64, error) {
	exoplanet, exists := exoplanets[id]
	if !exists {
		return 0, errors.New("exoplanet not found")
	}
	var gravity float64
	switch exoplanet.Type {
	case GasGiant:
		gravity = 0.5 / (exoplanet.Radius * exoplanet.Radius)
	case Terrestrial:
		if exoplanet.Mass == nil {
			return 0, errors.New("mass is required for Terrestrial type")
		}
		gravity = *exoplanet.Mass / (exoplanet.Radius * exoplanet.Radius)
	default:
		return 0, errors.New("invalid exoplanet type")
	}
	fuelCost := exoplanet.Distance / (gravity * gravity) * float64(crewCapacity)
	return fuelCost, nil
}
