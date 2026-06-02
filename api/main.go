package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sachithramanamperi/croptracetelemetry/config"
	"github.com/sachithramanamperi/croptracetelemetry/controllers"
	"github.com/sachithramanamperi/croptracetelemetry/database"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

func JSONLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		duration := time.Since(startTime)

		logrus.WithFields(logrus.Fields{
			"client_ip":  c.ClientIP(),
			"duration":   duration,
			"method":     c.Request.Method,
			"path":       c.Request.RequestURI,
			"status":     c.Writer.Status(),
			"user_agent": c.Request.UserAgent(),
		}).Info("Request processed")
	}
}

func main() {
	cfg := config.LoadConfig()

	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	err := database.InitDB(cfg)
	if err != nil {
		logrus.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.CloseDB()

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(JSONLogMiddleware())

	v1 := r.Group("/api/v1")
	{
		v1.POST("/telemetry", controllers.CreateTelemetry)
		v1.GET("/health", controllers.HealthCheck)
	}

	srv := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: r,
	}

	go func() {
		logrus.Infof("Server starting on port %s", cfg.AppPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatal("Server Shutdown:", err)
	}
	logrus.Info("Server exiting")
}
