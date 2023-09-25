package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Environments struct {
	APIPort    string
	DBHost     string
	DBPort     string
	DBDatabase string
	DBUser     string
	DBPassword string
}

func LoadEnvs() *Environments {
	if err := godotenv.Load(); err != nil {
		log.Print("Error loading .env file", err)
	}

	env := Environments{
		APIPort:    os.Getenv("PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBDatabase: os.Getenv("DB_DATABASE"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	}

	setEnvVariables(env)

	return &env
}

func setEnvVariables(env Environments) {
	os.Setenv("PORT", env.APIPort)
	os.Setenv("DB_HOST", env.DBHost)
	os.Setenv("DB_PORT", env.DBPort)
	os.Setenv("DB_DATABASE", env.DBDatabase)
	os.Setenv("DB_USER", env.DBUser)
	os.Setenv("DB_PASSWORD", env.DBPassword)
}
