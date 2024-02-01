package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func updateConfigWithEnvVariables() (*config, error) {
	// Load environment variables from `.env` file
	err := godotenv.Load("../../.env", "../../.env.development")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	maxOpenConnsStr := os.Getenv("DB_MAX_OPEN_CONNS")
	maxOpenConns, err := strconv.Atoi(maxOpenConnsStr)
	if err != nil {
		log.Fatal(err)
	}
	maxIdleConnsStr := os.Getenv("DB_MAX_IDLE_CONNS")
	maxIdleConns, err := strconv.Atoi(maxIdleConnsStr)
	if err != nil {
		log.Fatal(err)
	}

	var cfg config

	// Basic config
	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	// Database config
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("DATABASE_URL"), "PostgreSQL DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", maxOpenConns, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", maxIdleConns, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime,
		"db-max-idle-time",
		os.Getenv("DB_MAX_IDLE_TIME"),
		"PostgreSQL max connection idle time",
	)

	// Redis config
	flag.StringVar(&cfg.redisURL, "redis-url", os.Getenv("REDIS_URL"), "Redis URL")

	flag.Parse()

	return &cfg, nil
}
