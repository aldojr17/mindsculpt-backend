package repository

import (
	"mindsculpt/domain"

	"gorm.io/gorm"
)

type ImageGenerationRepository struct {
	db *gorm.DB
}

func NewImageGenerationRepository(db *gorm.DB) *ImageGenerationRepository {
	return &ImageGenerationRepository{
		db: db,
	}
}

func (r *ImageGenerationRepository) Create(payload domain.ImageGeneration) error {
	return r.db.Create(&payload).Error
}
