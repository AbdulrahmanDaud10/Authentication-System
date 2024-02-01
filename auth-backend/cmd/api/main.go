package main

import (
	"os"

	"github.com/AbdulrahmnaDaud10/auth-go/internal/jsonlog"
	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

const version = "1.0.0"

// `config` type to house all our app's configurations
type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
	redisURL string
}

// Main `application` type
type application struct {
	config      config
	logger      *jsonlog.Logger
	redisClient *redis.Client
}

func main() {
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	cfg, err := updateConfigWithEnvVariables()
	if err != nil {
		logger.PrintFatal(err, nil)
	}

	db, err := openDB(*cfg)

	if err != nil {
		logger.PrintFatal(err, nil)
	}

	defer db.Close()

	logger.PrintInfo("database connection pool established", nil)

	opt, err := redis.ParseURL(cfg.redisURL)
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	client := redis.NewClient(opt)

	logger.PrintInfo("redis connection pool established", nil)

	app := &application{
		config:      *cfg,
		logger:      logger,
		redisClient: client,
	}

	err = app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}
