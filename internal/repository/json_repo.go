package repository

import (
	"encoding/json"
	"fmt"
	"krushi-server/internal/models"
	"os"
	"sync"

	"github.com/google/uuid"
)

type JSONFarmerRepository struct {
	filePath string
	mu       sync.RWMutex
}

func NewJSONFarmerRepository(filePath string) *JSONFarmerRepository {
	return &JSONFarmerRepository{
		filePath: filePath,
	}
}

func (r *JSONFarmerRepository) Save(farmer *models.Farmer) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	farmers, err := r.loadFarmers()
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("could not load farmers: %w", err)
	}

	if farmer.FarmerID == "" {
		farmer.FarmerID = uuid.New().String()
	}

	farmers = append(farmers, *farmer)

	data, err := json.MarshalIndent(farmers, "", "  ")
	if err != nil {
		return fmt.Errorf("could not marshal farmers: %w", err)
	}

	return os.WriteFile(r.filePath, data, 0644)
}

func (r *JSONFarmerRepository) GetByID(id string) (*models.Farmer, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	farmers, err := r.loadFarmers()
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrFarmerNotFound
		}
		return nil, fmt.Errorf("could not load farmers: %w", err)
	}

	for _, f := range farmers {
		if f.FarmerID == id {
			return &f, nil
		}
	}

	return nil, ErrFarmerNotFound
}

func (r *JSONFarmerRepository) loadFarmers() ([]models.Farmer, error) {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}

	var farmers []models.Farmer
	if err := json.Unmarshal(data, &farmers); err != nil {
		return nil, fmt.Errorf("could not unmarshal farmers: %w", err)
	}

	return farmers, nil
}
