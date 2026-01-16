package main

import (
	"github.com/marketplace-go-backend/services/common/database"
	"gorm.io/gorm"
)

func NewStore() (*gorm.DB, error) {
	return database.NewStore()
}
