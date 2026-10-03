package repository

import "krushi-server/internal/models"

type FarmerRepository interface {
	Save(farmer *models.Farmer) error
	GetByID(id string) (*models.Farmer, error)
	ListAll() ([]*models.Farmer, error)
	GetByPhone(phone string) (*models.Farmer, error)
}
