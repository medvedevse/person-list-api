package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Config struct {
	DB     DBConfig
	Server ServerConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Url      string
}

type ServerConfig struct {
	Host string
	Port string
}

func InitConfig(l *zap.Logger) *Config {
	if err := godotenv.Load(); err != nil {
		l.Fatal("Went something wrong with .env config", zap.Error(err))
	}
	l.Info(".env config loaded")

	dbconfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PSWD"),
		Name:     os.Getenv("DB_NAME"),
	}
	l.Info("Database credentials loaded")
	serverConfig := ServerConfig{
		Host: os.Getenv("APP_HOST"),
		Port: os.Getenv("APP_PORT"),
	}
	l.Info("Server credentials loaded")
	dbconfig.Url = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbconfig.Host,
		dbconfig.Port,
		dbconfig.User,
		dbconfig.Password,
		dbconfig.Name,
	)
	l.Info("URL for database created")
	cfg := &Config{
		DB:     dbconfig,
		Server: serverConfig,
	}
	return cfg
}
