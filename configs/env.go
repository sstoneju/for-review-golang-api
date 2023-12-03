package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariableFromDot() {
	log.Println("working good?")
	env := os.Getenv("ENV")

	dotFile := func(env string) string {
		if env == "" {
			return "dev.env"
		}
		return env + ".env"
	}(env)

	log.Printf("Init environment variable from dot file thourgh lambda ENV. : %s", dotFile)

	err := godotenv.Overload(".env", dotFile)
	if err != nil {
		log.Println("Error loading .env file")
	}

	log.Println("Environment load complete!")
}
