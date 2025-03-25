package common

import (
	"log"
	"os"
)

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("⚠️ La variable de entorno %s no está definida.", key)
	}

	return value
}
