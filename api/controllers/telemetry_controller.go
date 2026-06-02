package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sachithramanamperi/croptracetelemetry/database"
	"github.com/sachithramanamperi/croptracetelemetry/models"
	"github.com/sirupsen/logrus"
)

func CreateTelemetry(c *gin.Context) {
	var req models.TelemetryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Warn("Invalid telemetry input data")
		
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "detail": err.Error()})
		return
	}

	record := models.TelemetryRecord{
		FacilityID:    req.FacilityID,
		Timestamp:     req.Timestamp,
		CropType:      req.CropType,
		WeightKg:      req.WeightKg,
		QualityRating: req.QualityRating,
	}

	if err := database.DB.Create(&record).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error":       err.Error(),
			"facility_id": req.FacilityID,
		}).Error("Failed to store telemetry data")
		
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store telemetry data"})
		return
	}

	logrus.WithFields(logrus.Fields{
		"facility_id": record.FacilityID,
		"crop_type":   record.CropType,
		"weight_kg":   record.WeightKg,
		"record_id":   record.ID,
	}).Info("Telemetry data stored")

	c.JSON(http.StatusCreated, record)
}
