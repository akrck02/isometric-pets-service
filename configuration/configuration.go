package configuration

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/withmandala/go-log"
)

var logger = log.New(os.Stderr)

type APIConfiguration struct {
	Ip         string
	Port       string
	Version    string
	ApiName    string
	Repository string

	CorsAccessControlAllowOrigin  string
	CorsAccessControlAllowMethods string
	CorsAccessControlAllowHeaders string
	CorsAccessControlMaxAge       string
}

func LoadConfiguration(path string) APIConfiguration {

	err := godotenv.Load(path)

	if nil != err {
		logger.Error("Error loading .env file")
	}

	var configuration = APIConfiguration{
		Ip:         os.Getenv("IP"),
		Port:       os.Getenv("PORT"),
		Version:    os.Getenv("VERSION"),
		ApiName:    os.Getenv("API_NAME"),
		Repository: os.Getenv("REPOSITORY"),

		CorsAccessControlAllowOrigin:  os.Getenv("CORS_ORIGIN"),
		CorsAccessControlAllowMethods: os.Getenv("CORS_METHODS"),
		CorsAccessControlAllowHeaders: os.Getenv("CORS_HEADERS"),
		CorsAccessControlMaxAge:       os.Getenv("CORS_MAX_AGE"),
	}

	checkCompulsoryVariables(configuration)
	return configuration
}

func checkCompulsoryVariables(Configuration APIConfiguration) {
	logger.Info("")
	logger.Info("Configuration variables")
	logger.Info()
	logger.Info("IP: " + Configuration.Ip)
	logger.Info("PORT: " + Configuration.Port)
	logger.Info("VERSION: " + Configuration.Version)
	logger.Info("API_NAME: " + Configuration.ApiName)
	logger.Info("REPOSITORY: " + Configuration.Repository)

	logger.Info()
	logger.Info("CORS_ORIGIN: " + Configuration.CorsAccessControlAllowOrigin)
	logger.Info("CORS_METHODS: " + Configuration.CorsAccessControlAllowMethods)
	logger.Info("CORS_HEADERS: " + Configuration.CorsAccessControlAllowHeaders)
	logger.Info("CORS_MAX_AGE: " + Configuration.CorsAccessControlMaxAge)
}

func (APIConfiguration) IsDevelopment() bool {
	return os.Getenv("ENV") == "development"
}
