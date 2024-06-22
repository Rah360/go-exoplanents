package repository

import (
	"errors"
	"exoplanet_microservice/models"

	"github.com/google/uuid"
)

type ExoplanetRepository interface {
	AddExoplanet(explanent models.Exoplanet) error
	ListExoplanets() ([]models.Exoplanet, error)
	GetExoplanetByID(id uuid.UUID) (models.Exoplanet, error)
	UpdateExoplanet(exoplanet models.Exoplanet) error
	DeleteExoplanet(id uuid.UUID) error
}

type InMemoryExoplanetRepository struct {
	exoplanets map[uuid.UUID]models.Exoplanet
}

func NewInMemoryExoplanentRepository() *InMemoryExoplanetRepository {
	return &InMemoryExoplanetRepository{
		exoplanets: make(map[uuid.UUID]models.Exoplanet),
	}
}

func (repo *InMemoryExoplanetRepository) AddExoplanet(exoplanet models.Exoplanet) error {
	repo.exoplanets[exoplanet.ID] = exoplanet
	return nil
}

func (repo *InMemoryExoplanetRepository) ListExoplanets() ([]models.Exoplanet, error) {
	planents := make([]models.Exoplanet, 0, len(repo.exoplanets))
	for _, exoplanent := range repo.exoplanets {
		planents = append(planents, exoplanent)
	}
	return planents, nil
}

func (repo *InMemoryExoplanetRepository) GetExoplanetByID(id uuid.UUID) (models.Exoplanet, error) {
	planet, exists := repo.exoplanets[id]
	if !exists {
		return models.Exoplanet{}, errors.New("exoplanet not found")
	}
	return planet, nil
}

func (repo *InMemoryExoplanetRepository) UpdateExoplanet(exoplanet models.Exoplanet) error {
	repo.exoplanets[exoplanet.ID] = exoplanet
	return nil
}

func (repo *InMemoryExoplanetRepository) DeleteExoplanet(id uuid.UUID) error {
	delete(repo.exoplanets, id)
	return nil
}
