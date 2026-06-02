package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TelemetryRecord struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	FacilityID    string    `gorm:"type:varchar(100);not null;index" json:"facility_id"`
	Timestamp     time.Time `gorm:"not null;index" json:"timestamp"`
	CropType      string    `gorm:"type:varchar(100);not null" json:"crop_type"`
	WeightKg      float64   `gorm:"not null" json:"weight_kg"`
	QualityRating int       `gorm:"not null" json:"quality_rating"`
	CreatedAt     time.Time `json:"created_at"`
}

func (record *TelemetryRecord) BeforeCreate(tx *gorm.DB) (err error) {
	if record.ID == uuid.Nil {
		record.ID = uuid.New()
	}
	return
}

type TelemetryRequest struct {
	FacilityID    string    `json:"facility_id" binding:"required"`
	Timestamp     time.Time `json:"timestamp" binding:"required"`
	CropType      string    `json:"crop_type" binding:"required"`
	WeightKg      float64   `json:"weight_kg" binding:"required,gt=0"`
	QualityRating int       `json:"quality_rating" binding:"required,gte=1,lte=10"`
}
