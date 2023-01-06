package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoLocalURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGODB_LOCAL_URI")
}

func EnvMongoStagingURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("MONGODB_STAGING_URI")
}

func EnvRedisURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("REDIS_LOCAL_URI")
}

func EnvRedisPassword() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("REDIS_LOCAL_PASSWORD")
}
