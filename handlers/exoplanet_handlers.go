package handlers

import (
	"encoding/json"
	"exoplanet_microservice/models"
	"exoplanet_microservice/services"
	"exoplanet_microservice/utils"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type ExoplanetHandler struct {
	service *services.ExoplanetService
}

func NewExoplanetHandler(service *services.ExoplanetService) *ExoplanetHandler {
	return &ExoplanetHandler{
		service: service,
	}
}

func (h *ExoplanetHandler) AddExoplanet(w http.ResponseWriter, r *http.Request) {
	var exoplanet models.Exoplanet
	if err := json.NewDecoder(r.Body).Decode(&exoplanet); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := exoplanet.Validate(); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	exoplanet.ID = uuid.New()
	if err := h.service.AddExoplanet(exoplanet); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.JSONResponse(w, http.StatusCreated, exoplanet)
}

func (h *ExoplanetHandler) ListExoplanets(w http.ResponseWriter, r *http.Request) {
	planets, err := h.service.ListExoplanets()
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.JSONResponse(w, http.StatusOK, planets)
}

func (h *ExoplanetHandler) GetExoplanetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	exoplanet, err := h.service.GetExoplanetByID(id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}
	utils.JSONResponse(w, http.StatusOK, exoplanet)
}

func (h *ExoplanetHandler) UpdateExoplanet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "invalid UUID")
		return
	}

	existingExoplanet, err := h.service.GetExoplanetByID(id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusNotFound, "exoplanet not found")
		return
	}

	var updateData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// Update only the provided fields
	if name, ok := updateData["name"].(string); ok {
		existingExoplanet.Name = name
	}
	if description, ok := updateData["description"].(string); ok {
		existingExoplanet.Description = description
	}
	if distance, ok := updateData["distance"].(float64); ok {
		existingExoplanet.Distance = int(distance)
	}
	if radius, ok := updateData["radius"].(float64); ok {
		existingExoplanet.Radius = radius
	}
	if mass, ok := updateData["mass"].(float64); ok {
		existingExoplanet.Mass = &mass
	}
	if exoplanetType, ok := updateData["type"].(string); ok {
		existingExoplanet.Type = models.ExoplanetType(exoplanetType)
	}

	if err := existingExoplanet.Validate(); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.UpdateExoplanet(existingExoplanet); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, existingExoplanet)
}

func (h *ExoplanetHandler) DeleteExoplanet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "planet id not found")
		return
	}
	if err := h.service.DeleteExoplanet(id); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "failed to delete planet")
		return
	}
	utils.JSONResponse(w, http.StatusNoContent, nil)
}

func (h *ExoplanetHandler) EstimateFuel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	crewCapacity := r.URL.Query().Get("crew")
	if crewCapacity == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "crew capacity is required")
		return
	}
	crew, err := strconv.Atoi(crewCapacity)
	if err != nil || crew <= 0 {
		utils.ErrorResponse(w, http.StatusBadRequest, "invalid crew capacity")
		return
	}
	fuelCost, err := h.service.EstimateFuel(id, crew)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.JSONResponse(w, http.StatusOK, map[string]float64{"fuel_cost": fuelCost})
}
