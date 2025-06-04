package http

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	_ "github.com/medvedevse/person-list-api/docs"
	"github.com/medvedevse/person-list-api/internal/repository/persistent"
	"github.com/medvedevse/person-list-api/internal/usecase"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(l *zap.Logger, router *gin.Engine, h *persistent.DBHandler) {
	// TODO: Probably it could be written better
	personHandler := usecase.Handler{H: *h}

	person := router.Group("/person")
	{
		person.GET("", personHandler.GetPersonList)
		person.DELETE("/:id", personHandler.DeletePerson)
		person.PUT("/:id", personHandler.UpdatePerson)
		person.POST("", personHandler.AddPerson)
	}

	router.GET("/", usecase.PreviewHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	l.Info("Routes successfully initialised")
}
