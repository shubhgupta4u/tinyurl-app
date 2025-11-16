package repository

import (
	"errors"
	"go-rest-api/config"
	"go-rest-api/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Create Link
func CreateLink(link *models.Link) error {
	return config.DB.Create(link).Error
}

// Get by ShortCode (active and not expired)
func GetLinkByShortCode(shortCode string) (*models.Link, error) {
	var link models.Link
	now := time.Now()
	result := config.DB.Preload("Organization").Preload("User").Where("short_code = ? AND is_active = TRUE AND (expires_at IS NULL OR expires_at > ?)", shortCode, now).First(&link)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &link, result.Error
}

// Delete by ShortCode
func DeleteLinkByShortCode(shortCode string) error {
	return config.DB.Where("short_code = ?", shortCode).Delete(&models.Link{}).Error
}

// Search by tags, user or org, filter by expiration
func SearchLinks(userID, orgID *uuid.UUID, tag string) ([]models.Link, error) {
	var links []models.Link
	query := config.DB.Model(&models.Link{}).Where("is_active = TRUE")
	if userID != nil {
		query = query.Where("user_id = ?", *userID)
	}
	if orgID != nil {
		query = query.Where("org_id = ?", *orgID)
	}
	if tag != "" {
		query = query.Where("? = ANY(tags)", tag)
	}
	query = query.Preload("Organization").Preload("User").Where("(expires_at IS NULL OR expires_at > ?)", time.Now())
	err := query.Find(&links).Error
	return links, err
}

// Update entire link
func UpdateLink(link *models.Link) error {
	return config.DB.Save(link).Error
}

// Patch link (partial update)
func PatchLink(id uint64, updates map[string]interface{}) error {
	return config.DB.Model(&models.Link{}).Where("id = ?", id).Updates(updates).Error
}
