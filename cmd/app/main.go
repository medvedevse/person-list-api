package main

import (
	"fmt"

	"github.com/medvedevse/person-list-api/config"
	"github.com/medvedevse/person-list-api/internal/controller/http"
	"github.com/medvedevse/person-list-api/internal/controller/http/middleware"
	"github.com/medvedevse/person-list-api/internal/repository/persistent"
	"github.com/medvedevse/person-list-api/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @title           Person List API
// @version         1.0
// @description     API Server for the PersonList Application

func main() {
	logger := logger.InitLogger()
	defer logger.Sync()

	cfg := config.InitConfig(logger)
	db := persistent.Connect(logger, cfg.DB.Url)
	h := &persistent.DBHandler{DB: db}

	router := gin.Default()
	router.Use(middleware.LoggerMiddleware(logger))
	http.InitRoutes(logger, router, h)

	router.Run(fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port))
	logger.Info("Server is running", zap.String("port", cfg.Server.Port))
}
