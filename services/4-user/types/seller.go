package types

import (
	"time"

	"gorm.io/gorm"
)

type Seller struct {
	ID              string         `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	BuyerID         string         `gorm:"type:uuid;not null" json:"buyerId"`
	FullName        string         `json:"fullName"`
	Bio             string         `json:"bio"`
	Email           string         `gorm:"uniqueIndex;not null" json:"email"`
	RatingsCount    int64          `gorm:"default:0" json:"ratingsCount"`
	RatingSum       int64          `gorm:"default:0" json:"ratingSum"`
	StripeAccountID string         `json:"stripeAccountID"`
	AccountBalance  uint64         `gorm:"default:0" json:"accountBalance"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
