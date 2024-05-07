package repositories

import (
	"OnlineStoreBackend/models"

	"gorm.io/gorm"
)

type RepositoryUser struct {
	DB *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *RepositoryUser {
	return &RepositoryUser{DB: db}
}

func (repository *RepositoryUser) ReadByID(modelUser *models.Users, UserID uint64) error {
	return repository.DB.First(modelUser, UserID).Error
}

func (repository *RepositoryUser) ReadAll(modelUsers *[]models.Users) error {
	return repository.DB.Find(modelUsers).Error
}
