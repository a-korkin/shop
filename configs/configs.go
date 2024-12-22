package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	godotenv.Load()
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("failed to load env variable: %s", key)
	}
	return value
}

func GetWebApiPort() string {
	return getEnv("WEB_API_PORT")
}
