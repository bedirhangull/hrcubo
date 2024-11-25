package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		DB     *DB
		Server *Server
		JWT    *JWT
		Redis  *Redis
	}

	DB struct {
		Connection string
		Host       string
		Port       string
		User       string
		Name       string
		Password   string
	}

	Server struct {
		Url  string
		Port string
	}

	JWT struct {
		Secret string
		Expire int
	}

	Redis struct {
		Host     string
		Port     string
		Password string
	}
)

func New() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		currentDir, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		envPath := filepath.Join(currentDir, ".env")
		log.Println("Loading env file from: ", envPath)

		err = godotenv.Load(envPath)
		if err != nil {
			return nil, err
		}
	}

	server := &Server{
		Url:  os.Getenv("HTTP_URL"),
		Port: os.Getenv("HTTP_PORT"),
	}

	db := &DB{
		Connection: os.Getenv("DB_CONNECTION"),
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		Name:       os.Getenv("DB_NAME"),
		User:       os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASSWORD"),
	}

	jwt := &JWT{
		Secret: os.Getenv("JWT_SECRET"),
		Expire: 60,
	}

	redis := &Redis{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}

	return &Container{
		DB:     db,
		Server: server,
		JWT:    jwt,
		Redis:  redis,
	}, nil
}
