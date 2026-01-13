package service

import "gorm.io/gorm"

type SellerService struct {
	db *gorm.DB
}

func NewSellerService(db *gorm.DB) *SellerService {
	return &SellerService{
		db: db,
	}
}

// TODO: Implement seller service methods
