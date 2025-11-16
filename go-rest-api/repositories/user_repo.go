package repository

import (
	"errors"
	"go-rest-api/config"
	"go-rest-api/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Create user
func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

// Get by ID
func GetUserByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	result := config.DB.Preload("Organization").First(&user, "id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, result.Error
}

// Get by Email
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := config.DB.Preload("Organization").First(&user, "email = ?", email)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, result.Error
}

// Update user
func UpdateUser(user *models.User) error {
	return config.DB.Save(user).Error
}

// Patch user (partial update)
func PatchUser(id uuid.UUID, updates map[string]interface{}) error {
	return config.DB.Model(&models.User{}).Where("id = ?", id).Updates(updates).Error
}

// Delete user
func DeleteUserByID(id uuid.UUID) error {
	return config.DB.Delete(&models.User{}, "id = ?", id).Error
}

// List users by org
func ListUsersByOrg(orgID uuid.UUID) ([]models.User, error) {
	var users []models.User
	err := config.DB.Preload("Organization").Where("org_id = ?", orgID).Find(&users).Error
	return users, err
}
