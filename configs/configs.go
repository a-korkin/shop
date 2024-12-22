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

func GetDBConnection() string {
	return getEnv("DB_CONN")
}

func GetWebApiPort() string {
	return getEnv("WEB_API_PORT")
}

func GetGrpcPort() string {
	return getEnv("GRPC_PORT")
}
