package application

import (
	"os"
	"strconv"
)

type config struct {
	RedisAddress string
	ServerPort   uint16
}

func LoadConfig() config {
	cfg := config{
		RedisAddress: "localhost:6379",
		ServerPort:   8080,
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
