package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port        string
	Db          db
	Environment string
}

type db struct {
	Name     string
	Port     string
	User     string
	Password string
}

func LoadConfig() Config {
	cfg := Config{}

	mustMapEnv(&cfg.Db.Name, "DB_NAME")
	mustMapEnv(&cfg.Db.Password, "DB_PASS")
	mustMapEnv(&cfg.Db.User, "DB_USER")
	mustMapEnv(&cfg.Db.Port, "DB_PORT")
	mustMapEnv(&cfg.Port, "PORT")

	if os.Getenv("ENV") != "" {
		cfg.Environment = os.Getenv("ENV")
	} else {
		cfg.Environment = "development"
	}
	fmt.Println(cfg)
	return cfg
}

func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		panic(fmt.Sprintf("environment variable %s not set", envKey))
	}
	*target = v
}
