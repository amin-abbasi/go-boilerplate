package configs

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const TIME_OUT_DURATION = 10 * time.Second

// load/read the .env file and return the value of the key
func GetEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}
