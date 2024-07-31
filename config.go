package main

import (
	"fmt"
	"os"
)

type Config struct {
	Port      string
	DBUser    string
	DBPass    string
	DBHost    string
	DBAddress string
	DBPort    string
	DBName    string
	JWTSecret string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		Port:      getEnv("PORT", "8080"),
		DBUser:    getEnv("DB_USER", "root"),
		DBPass:    getEnv("DB_PASS", "password"),
		DBAddress: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "localhost"), getEnv("DB_PORT", "3306")),
		DBName:    getEnv("DB_NAME", "go_backend"),
		JWTSecret: getEnv("JWT_SECRET", "randomjwtsecret"),
	}

}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
