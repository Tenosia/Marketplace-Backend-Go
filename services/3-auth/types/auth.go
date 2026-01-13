package types

import (
	"time"

	"gorm.io/gorm"
)

type Auth struct {
	ID                     string         `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Username               string         `gorm:"uniqueIndex;not null" json:"username"`
	Email                  string         `gorm:"uniqueIndex;not null" json:"email"`
	Password               string         `gorm:"not null" json:"-"`
	ProfilePublicID        string         `json:"profilePublicID"`
	Country                string         `json:"country"`
	ProfilePicture         string         `json:"profilePicture"`
	EmailVerified          bool           `gorm:"default:false" json:"emailVerified"`
	EmailVerificationToken string         `json:"-"`
	PasswordResetToken     string         `json:"-"`
	PasswordResetExpires   *time.Time     `json:"-"`
	CreatedAt              time.Time      `json:"createdAt"`
	UpdatedAt              time.Time      `json:"updatedAt"`
	DeletedAt              gorm.DeletedAt `gorm:"index" json:"-"`
}
