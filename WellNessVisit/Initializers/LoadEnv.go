package initializers

import (
	"github.com/joho/godotenv"
)

func LoadEnvVariables() error {
	return godotenv.Load(".env")
}
