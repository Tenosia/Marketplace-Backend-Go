package service

import "gorm.io/gorm"

type BuyerService struct {
	db *gorm.DB
}

func NewBuyerService(db *gorm.DB) *BuyerService {
	return &BuyerService{
		db: db,
	}
}

// TODO: Implement buyer service methods
