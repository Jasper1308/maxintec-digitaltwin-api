package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("[Aviso]: Arquivo .env não encontrado, usando variáveis de ambiente do sistema.")
	}

	return &Config{
		DBUser:     os.Getenv("DATABASEUSER"),
		DBPassword: os.Getenv("DATABASEPASSWORD"),
		DBHost:     os.Getenv("DATABASEHOST"),
	}
}