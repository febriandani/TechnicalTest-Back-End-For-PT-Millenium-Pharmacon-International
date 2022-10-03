package controller

import (
	"technical-test-02-10-22/models"

	"gorm.io/gorm"
)

type Repository interface {
	Create(file models.File) (models.File, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(file models.File) (models.File, error) {
	err := r.db.Create(&file).Error
	if err != nil {
		return models.File{}, err
	}

	return models.File{}, nil
}
