package pkg

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	BLIZZARD_CLIENT_SECRET string
	BLIZZARD_CLIENT_ID     string
}

var Envs = config()

func config() *Config {
	return &Config{
		BLIZZARD_CLIENT_SECRET: getEnv("BLIZZARD_CLIENT_SECRET", ""),
		BLIZZARD_CLIENT_ID:     getEnv("BLIZZARD_CLIENT_ID", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
