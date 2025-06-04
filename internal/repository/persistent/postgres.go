package persistent

import (
	"log"
	"os"
	"time"

	"github.com/medvedevse/person-list-api/internal/entity"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBHandler struct {
	DB *gorm.DB
}

func initGormLogger() logger.Interface {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Error,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)
	return newLogger
}

func Connect(l *zap.Logger, dbUrl string) *gorm.DB {
	dbLogger := initGormLogger()
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		Logger: dbLogger,
	})
	if err != nil {
		l.Fatal("Error connecting to database", zap.Error(err))
	}

	l.Info("Launching database automigration")
	db.AutoMigrate(&entity.Person{})
	l.Info("Connection to the database established")
	return db
}
