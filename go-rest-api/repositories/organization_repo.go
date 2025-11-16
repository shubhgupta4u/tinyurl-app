package repository

import (
	"errors"
	"go-rest-api/config"
	"go-rest-api/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Create organization
func CreateOrganization(org *models.Organization) error {
	return config.DB.Create(org).Error
}

// Get by ID
func GetOrganizationByID(id uuid.UUID) (*models.Organization, error) {
	var org models.Organization
	result := config.DB.First(&org, "id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &org, result.Error
}

// Update organization
func UpdateOrganization(org *models.Organization) error {
	return config.DB.Save(org).Error
}

// Delete organization
func DeleteOrganizationByID(id uuid.UUID) error {
	return config.DB.Delete(&models.Organization{}, "id = ?", id).Error
}

// List all organizations
func ListOrganizations() ([]models.Organization, error) {
	var orgs []models.Organization
	err := config.DB.Find(&orgs).Error
	return orgs, err
}
