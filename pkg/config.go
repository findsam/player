package pkg

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	BLIZZARD_CLIENT_SECRET,
	BLIZZARD_CLIENT_ID,
	DB_PWD,
	DB_USER,
	DB_NAME string
}

var Envs = config()

func config() *Config {
	return &Config{
		BLIZZARD_CLIENT_SECRET: getEnv("BLIZZARD_CLIENT_SECRET", ""),
		BLIZZARD_CLIENT_ID:     getEnv("BLIZZARD_CLIENT_ID", ""),
		DB_PWD: getEnv("DB_PWD", ""),
		DB_USER: getEnv("DB_USER", ""),
		DB_NAME: getEnv("DB_NAME", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
