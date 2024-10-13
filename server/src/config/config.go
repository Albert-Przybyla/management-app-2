package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	cfg  *Config
	once sync.Once
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPass     string
	DBName     string
	JWTSecret  string
	UsersTable string
	Port       string
}

func LoadConfigSingleton() (*Config, error) {
	var err error
	once.Do(func() {
		err = loadConfig()
	})
	return cfg, err
}

func loadConfig() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("failed to load .env file: %w", err)
	}

	cfg = &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPass:     os.Getenv("DB_PASS"),
		DBName:     os.Getenv("DB_DB"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
		UsersTable: os.Getenv("USERS_TABLE"),
		Port:       os.Getenv("PORT"),
	}

	missingEnvVars := checkEnvVars(cfg)
	if len(missingEnvVars) > 0 {
		return fmt.Errorf("missing environment variables: %v", missingEnvVars)
	}

	return nil
}

func checkEnvVars(config *Config) []string {
	var missing []string
	if config.DBHost == "" {
		missing = append(missing, "DB_HOST")
	}
	if config.DBPort == "" {
		missing = append(missing, "DB_PORT")
	}
	if config.DBUser == "" {
		missing = append(missing, "DB_USER")
	}
	if config.DBPass == "" {
		missing = append(missing, "DB_PASS")
	}
	if config.DBName == "" {
		missing = append(missing, "DB_DB")
	}
	if config.JWTSecret == "" {
		missing = append(missing, "JWT_SECRET")
	}
	if config.UsersTable == "" {
		missing = append(missing, "USERS_TABLE")
	}
	if config.Port == "" {
		missing = append(missing, "PORT")
	}
	return missing
}
