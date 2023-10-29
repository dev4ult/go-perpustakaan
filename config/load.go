package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	DB_USER string
	DB_PASS string
	DB_HOST	string
	DB_PORT	int
	DB_NAME string
}

type ServerConfig struct {
	SERVER_PORT int
	SIGN_KEY string
	CLOUD_URL string
	REFRESH_KEY string
	HGF_TOKEN string
	MT_SERVER_KEY string
}

func LoadDBConfig() DatabaseConfig {
	godotenv.Load(".env")

	DB_PORT, err := strconv.Atoi(os.Getenv("DB_PORT"))
	
	if err != nil {
		panic(err)
	}

	return DatabaseConfig {
		DB_USER: os.Getenv("DB_USER"),
		DB_PASS: os.Getenv("DB_PASS"),
		DB_HOST: os.Getenv("DB_HOST"),
		DB_PORT: DB_PORT,
		DB_NAME: os.Getenv("DB_NAME"),
	}
}

func LoadServerConfig() ServerConfig {
	godotenv.Load(".env")

	SERVER_PORT, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	
	if err != nil {
		SERVER_PORT = 8000
	}

	return ServerConfig{
		SERVER_PORT: SERVER_PORT,
		SIGN_KEY: os.Getenv("SIGN_KEY"),
		CLOUD_URL: os.Getenv("CLOUDINARY_URL"),
		REFRESH_KEY: os.Getenv("REFRESH_KEY"),
		HGF_TOKEN: os.Getenv("HGF_TOKEN"),
		MT_SERVER_KEY: os.Getenv("MT_SERVER_KEY"),
	}
}

