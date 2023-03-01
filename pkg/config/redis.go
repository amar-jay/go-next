package config

import (
	"os"
)

// PostgresConfig is the config for postgres
type RedisConfig struct {
	Address string `env:"REDIS_ADDRESS"`
}

// GetPostgresConfig returns the postgres config
func GetRedisConfig() RedisConfig {
	return RedisConfig{
		Address: os.Getenv("REDIS_ADDRESS"),
	}
}
