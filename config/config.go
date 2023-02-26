package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
)

type Database struct {
	DB_NAME     string
	DB_USERNAME string
	DB_HOST     string
	DB_PASSWORD string
	DB_PORT     int
}

type Config struct {
	Database  Database
	DebugMode bool
}

func New() *Config {
	return &Config{
		Database: Database{
			DB_NAME:     getEnv("DB_NAME", ""),
			DB_USERNAME: getEnv("DB_USERNAME", ""),
			DB_HOST:     getEnv("DB_HOST", ""),
			DB_PASSWORD: getEnv("DB_PASSWORD", ""),
			DB_PORT:     getEnvAsInt("DB_PORT", 5432),
		},
		DebugMode: getEnvAsBool("DEBUG_MODE", true),
	}
}

func init() {
	Error := godotenv.Load()
	if Error != nil {
		log.Print(".env file no found")
	}
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
