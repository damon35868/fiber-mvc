package config

import (
	"os"
	"runtime"
	"strconv"

	"github.com/gofiber/storage/redis/v3"
)

func NewRedis() *redis.Storage {
	port, _ := strconv.Atoi(os.Getenv("CACHE_PORT"))
	db, _ := strconv.Atoi(os.Getenv("CACHE_DB"))

	return redis.New(redis.Config{
		Database:   db,
		Port:       port,
		Reset:      false,
		TLSConfig:  nil,
		ClientName: os.Getenv("CACHE_PREFIX"),
		Host:       os.Getenv("CACHE_HOST"),
		Username:   os.Getenv("CACHE_USERNAME"),
		Password:   os.Getenv("CACHE_PASSWORD"),
		PoolSize:   10 * runtime.GOMAXPROCS(0),
	})

}
