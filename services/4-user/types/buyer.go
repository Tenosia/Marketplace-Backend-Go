package types

import (
	"time"

	"gorm.io/gorm"
)

type Buyer struct {
	ID             string         `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Username       string         `gorm:"uniqueIndex;not null" json:"username"`
	Email          string         `gorm:"uniqueIndex;not null" json:"email"`
	Country        string         `json:"country"`
	ProfilePicture string         `json:"profilePicture"`
	IsSeller       bool           `gorm:"default:false" json:"isSeller"`
	StripeAccountId string         `json:"stripeAccountId"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}
