package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Env         string
	StoragePath string
	TokenTTL    time.Duration
	GRPCConfig
}

type GRPCConfig struct {
	Port    string
	Timeout time.Duration
}

func MustLoad() *Config {
	const baseError = "error loading .env file"
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("%s: %s", baseError, err.Error())
	}

	var cfg Config

	// ENV
	cfg.Env = os.Getenv("ENV")
	if cfg.Env == "" {
		log.Fatalf("%s: ENV is required", baseError)
	}
	if cfg.Env != "local" && cfg.Env != "dev" && cfg.Env != "prod" {
		log.Fatalf("%s: invalid ENV value. Must be 'local', 'dev' or 'prod'", baseError)
	}

	// STORAGE_PATH
	cfg.StoragePath = os.Getenv("STORAGE_PATH")
	if cfg.Env == "" {
		log.Fatalf("%s: STORAGE_PATH is required", baseError)
	}

	// TOKEN_TTL
	TTLString := os.Getenv("TOKEN_TTL")
	if TTLString == "" {
		log.Fatalf("%s: TOKEN_TTL is required", baseError)
	}
	TTLInt, err := strconv.Atoi(TTLString)
	if err != nil {
		log.Fatalf("%s: invalid TOKEN_TTL value. Must be int type", baseError)
	}
	cfg.TokenTTL = time.Duration(TTLInt) * time.Hour

	// GRPC_PORT
	cfg.GRPCConfig.Port = os.Getenv("GRPC_PORT")
	if cfg.Env == "" {
		log.Fatalf("%s: GRPC_PORT is required", baseError)
	}

	// GRPC_TIMEOUT
	timeOutString := os.Getenv("GRPC_TIMEOUT")
	if timeOutString == "" {
		log.Fatalf("%s: GRPC_TIMEOUT is required", baseError)
	}
	timeOutInt, err := strconv.Atoi(timeOutString)
	if err != nil {
		log.Fatalf("%s: invalid GRPC_TIMEOUT value. Must be int type", baseError)
	}
	cfg.GRPCConfig.Timeout = time.Duration(timeOutInt) * time.Second

	return &cfg
}
