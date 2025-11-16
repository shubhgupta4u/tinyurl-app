package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Link struct {
	ID          uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      uuid.UUID      `gorm:"type:uuid;not null" json:"user_id"`
	OrgID       uuid.UUID      `gorm:"type:uuid;not null" json:"org_id"`
	ShortCode   string         `gorm:"not null;uniqueIndex:idx_domain_shortcode" json:"short_code"`
	TargetURL   string         `gorm:"not null" json:"target_url" validate:"required,url"`
	Domain      string         `gorm:"not null;default:'tiny.example.com'" json:"domain"`
	CustomAlias bool           `json:"custom_alias"`
	Title       string         `json:"title"`
	Tags        pq.StringArray `gorm:"type:text[]" json:"tags"`
	CampaignID  *uuid.UUID     `gorm:"type:uuid" json:"campaign_id,omitempty"`
	ExpiresAt   *time.Time     `json:"expires_at,omitempty"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`

	// GORM relation
	Organization Organization `gorm:"foreignKey:OrgID" json:"organization,omitempty"`
	User         User         `gorm:"foreignKey:UserID" json:"user,omitempty"`
}
