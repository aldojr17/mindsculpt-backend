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

func (r *ImageGenerationRepository) GetByUUID(uuid string) (*domain.ImageGeneration, error) {
	data := new(domain.ImageGeneration)
	if err := r.db.Where("id", uuid).First(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
