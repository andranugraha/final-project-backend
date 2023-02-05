package config

import (
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

type dbConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func getENV(key, defaultVal string) string {
	env := os.Getenv(key)
	if env == "" {
		return defaultVal
	}
	return env
}

func getArrayENV(key string) []string {
	env := os.Getenv(key)
	if env == "" {
		return []string{}
	}

	return strings.Split(env, ",")
}

var (
	ENV      = getENV("APP_ENV", "test")
	AppName  = "sea-labs-library"
	DBConfig = dbConfig{
		Host:     getENV("DB_HOST", "localhost"),
		User:     getENV("DB_USER", ""),
		Password: getENV("DB_PASSWORD", ""),
		DBName:   getENV("DB_NAME", ""),
		Port:     getENV("DB_PORT", "5432"),
	}
	Secret        = getENV("SECRET", "secret")
	CloudinaryUrl = getENV("CLOUDINARY_URL", "cloudinary://")
	AllowOrigins  = getArrayENV("ALLOW_ORIGINS")
)
