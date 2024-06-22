package services

import (
	"exoplanet_microservice/models"
	"exoplanet_microservice/repository"

	"github.com/google/uuid"
)

type ExoplanetService struct {
	repo repository.ExoplanetRepository
}

func NewExoplanetService(repo repository.ExoplanetRepository) *ExoplanetService {
	return &ExoplanetService{
		repo: repo,
	}
}

func (s *ExoplanetService) AddExoplanet(exoplanet models.Exoplanet) error {
	return s.repo.AddExoplanet(exoplanet)
}

func (s *ExoplanetService) ListExoplanets() ([]models.Exoplanet, error) {
	return s.repo.ListExoplanets()
}

func (s *ExoplanetService) GetExoplanetByID(id uuid.UUID) (models.Exoplanet, error) {
	return s.repo.GetExoplanetByID(id)
}

func (s *ExoplanetService) UpdateExoplanet(exoplanet models.Exoplanet) error {
	return s.repo.UpdateExoplanet(exoplanet)
}

func (s *ExoplanetService) DeleteExoplanet(id uuid.UUID) error {
	return s.repo.DeleteExoplanet(id)
}

func (s *ExoplanetService) EstimateFuel(id uuid.UUID, crewCapacity int) (float64, error) {
	exoplanet, err := s.repo.GetExoplanetByID(id)
	if err != nil {
		return 0, err
	}
	return exoplanet.FuelEstimation(crewCapacity), nil
}
