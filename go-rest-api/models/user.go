package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	OrgID        uuid.UUID `gorm:"type:uuid;not null" json:"org_id"`
	Name         string    `gorm:"not null" json:"name"`
	Email        string    `gorm:"unique;not null" json:"email"`
	Mobile       string    `gorm:"unique" json:"mobile,omitempty"`
	PasswordHash string    `gorm:"not null" json:"-"`
	IsActive     bool      `gorm:"default:true" json:"is_active"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`

	// GORM relation
	Organization Organization `gorm:"foreignKey:OrgID" json:"organization,omitempty"`
}
