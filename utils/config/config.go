package config

import (
	"os"
)
const (
  // AppName is the name of the app
  AppName = "go-auth"
  production = "production"
)


type Config struct {
  Pepper    string        `env:"PEPPER"`
  Env       string        `env:"ENV"`
  FromEmail string        `env:"EMAIL_FROM"`
  HashKey   string        `env:"HASH_KEY"`
  Port      int           `env:"PORT"`
  JWTSecret string        `env:"JWT_SECRET"`
  Mailgun   MailgunConfig `json:"mailgun"`
  Postgres  PostgresConfig `json:"postgres"`
  Redis	    RedisConfig	  `json:"redis"`
}

// Check if it is in production
func (c Config) IsProduction() bool {
  return c.Env == production
}

func GetConfig() Config {
  return Config{
    Pepper:  os.Getenv("PEPPER"),
    Env: os.Getenv("ENV"),
    Mailgun: GetMailgunConfig(),
    Postgres: GetPostgresConfig(),
    FromEmail: os.Getenv("EMAIL_FROM"),
    Port: getPort("PORT"),
    HashKey: os.Getenv("HASH_KEY"),
    JWTSecret: os.Getenv("JWT_SECRET"),
  }
}

type Key string
func GinContextKey() Key {
  return Key("GinContextKey")
}
