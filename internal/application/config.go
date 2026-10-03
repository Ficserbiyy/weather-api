package application

import (
	"log"
	"os"
	"strconv"
)

type config struct {
	RedisAddress string
	ServerPort   uint16
	APIKey       string
}

const (
	envAPIKey = "WEATHER_API_KEY"
)

func LoadConfig() config {
	apiKey := os.Getenv(envAPIKey)
	if apiKey == "" {
		log.Fatalf("environment variable is not set: %s", envAPIKey)
	}

	cfg := config{
		RedisAddress: "localhost:6379",
		ServerPort:   8080,
		APIKey:       apiKey,
	}

	if redisAddr, ok := os.LookupEnv("REDIS_ADDR"); ok {
		cfg.RedisAddress = redisAddr
	}

	if serverPort, ok := os.LookupEnv("SERVER_PORT"); ok {
		if port, err := strconv.ParseUint(serverPort, 10, 16); err == nil {
			cfg.ServerPort = uint16(port)
		}
	}

	return cfg
}
