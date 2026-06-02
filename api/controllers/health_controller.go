package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sachithramanamperi/croptracetelemetry/database"
	"github.com/sirupsen/logrus"
)

func HealthCheck(c *gin.Context) {
	if database.DB == nil {
		logrus.Warn("Health check failed: database not initialized")
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "unhealthy", "database": "disconnected"})
		return
	}

	sqlDB, err := database.DB.DB()
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err.Error()}).Warn("Health check failed: could not get sql.DB")
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "unhealthy", "database": "disconnected"})
		return
	}

	if err := sqlDB.Ping(); err != nil {
		logrus.WithFields(logrus.Fields{"error": err.Error()}).Warn("Health check failed: database ping failed")
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "unhealthy", "database": "disconnected"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "healthy",
		"database": "connected",
	})
}
