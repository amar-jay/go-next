package config

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresConfig is the config for postgres
type PostgresConfig struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     int    `env:"POSTGRES_PORT"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Database string `env:"POSTGRES_DB"`
}

// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// Dialect returns the dialect for postgres
func (c *PostgresConfig) Config() *gorm.Config {
	return &gorm.Config{}
}

// GetConnectionInfo returns the connection info for postgres
func (c *PostgresConfig) GetConnectionInfo() gorm.Dialector {
	fmt.Println(c.Host, c.Port, c.User, c.Password, c.Database)
	dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Database)
	return postgres.Open(dns)
}

// GetPostgresConfig returns the postgres config
func GetPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     getPort("POSTGRES_PORT", 5432),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
	}
}

func getPort(input string, def uint) int {
	port, err := strconv.Atoi(os.Getenv(input))
	if err != nil {
		fmt.Printf("error in parsing %s. Setting to default %d", input, def)
		panic(err)
	}
	if port < 0 {
		panic(fmt.Errorf("invalid port number: %d", port))
	}

	return port
}
