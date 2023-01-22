package initializers

import (
	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	godotenv.Load(".env")
}
