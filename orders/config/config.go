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
	ProjectId string
}

func LoadConfig() Config {
	cfg := Config{}
	mustMapEnv(&cfg.Db.ProjectId, "FIRESTORE_PROJ_ID")
	mustMapEnv(&cfg.Port, "PORT")
	if os.Getenv("ENV") != "" {
		cfg.Environment = os.Getenv("ENV")
	} else {
		cfg.Environment = "development"
	}
	return cfg
}

func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		panic(fmt.Sprintf("environment variable %s not set", envKey))
	}
	*target = v
}
