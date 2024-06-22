package models

import (
	"errors"
	"math"

	"github.com/google/uuid"
)

type ExoplanetType string

const (
	GasGiant    ExoplanetType = "gasGiant"
	Terrestrial ExoplanetType = "terrestrial"
)

type Exoplanet struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Distance    int           `json:"distance"`
	Radius      float64       `json:"radius"`
	Mass        *float64      `json:"mass,omitempty"`
	Type        ExoplanetType `json:"type"`
}

func (e *Exoplanet) Validate() error {
	if e.Distance <= 10 || e.Distance >= 1000 {
		return errors.New("distance must be between 10 and 1000 ligth years")
	}
	if e.Radius <= 0.1 || e.Radius >= 10 {
		return errors.New("radius must be between 0.1 and 10 Earth-radius units")
	}

	if e.Type == Terrestrial {
		if e.Mass == nil || *e.Mass <= 0.1 || *e.Mass >= 10 {
			return errors.New("mass must be provided and between 0.1 and 10 Earth-mass units for terrestrial planets")
		}
	}
	return nil
}

func (e *Exoplanet) Gravity() float64 {
	if e.Type == GasGiant {
		return 0.5 / math.Pow(e.Radius, 2)
	} else if e.Type == Terrestrial {
		return *e.Mass / math.Pow(e.Radius, 2)
	}
	return 0
}

func (e *Exoplanet) FuelEstimation(crewCapacity int) float64 {
	gravity := e.Gravity()
	return float64(e.Distance) / math.Pow(gravity, 2) * float64(crewCapacity)
}
