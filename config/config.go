package config

import (
	"log"

	"github.com/joho/godotenv"
)

// Config describes configuration.
type Config struct {
	Port               string `default:"3000"`
	Service            string `required:"true"`
	Endpoint           string `required:"true"`
	Region             string `required:"true"`
	AwsAccessKeyID     string `envconfig:"AWS_ACCESS_KEY_ID" required:"true"`
	AwsSecretAccessKey string `envconfig:"AWS_SECRET_ACCESS_KEY" required:"true"`
	AwsSessionToken    string `envconfig:"AWS_SESSION_TOKEN"`
}

func init() {
	log.Println("Loading .env")
	err := godotenv.Load()
	if err != nil {
		log.Print("No .env file, running in production mode.")
	}
}
