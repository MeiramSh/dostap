package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port     int
	Host     string
	Postgres PostgresConfig
	Email    EmailConfig
}
type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type EmailConfig struct {
	Host     string
	Username string
	Password string
	Identity string
	Addr     string
}

func NewConfig() *Config {
	port, err := strconv.Atoi(getEnv("SERVER_PORT", "8080"))
	if err != nil {
		port = 8080
	}
	postgresPort, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		postgresPort = 5432
	}

	cfg := &Config{
		Port: port,
		Host: getEnv("SERVER_HOST", "localhost"),
		Postgres: PostgresConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     postgresPort,
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "password"),
			DBName:   getEnv("DB_NAME", "drev"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		Email: EmailConfig{
			Host:     getEnv("EMAIL_HOST", "smtp.example.com"),
			Username: getEnv("EMAIL_USERNAME", "username@gmail.com"),
			Password: getEnv("EMAIL_PASSWORD", "password"),
			Identity: getEnv("EMAIL_IDENTITY", ""),
			Addr:     getEnv("EMAIL_ADDR", "smtp.example.com:587"),
		},
	}

	return cfg
}

func getEnv(key, defaultValue string) string {
	value, found := os.LookupEnv(key)
	if !found {
		return defaultValue
	}
	return value
}

func (cfg *Config) Addr() string {
	return fmt.Sprintf("%s:%v", cfg.Host, cfg.Port)
}
func (pc *PostgresConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		pc.Host, pc.Port, pc.User, pc.Password, pc.DBName, pc.SSLMode)
}
