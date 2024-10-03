package Env

import (
	"os"
	"strings"
)

type Env struct {
	GoogleClientId      string
	GoogleClientSecret  string
	DatabaseUrl         string
	AllowedRedirectUrls []string
	Port                string
	DbHost              string
	DbPort              string
	DbUser              string
	DbPassword          string
	DbName              string
}

func ReadEnvVar(key string, fallback string) string {
	val := os.Getenv(key)

	if val == "" {
		return fallback
	}

	return val
}

func NewEnv() *Env {
	return &Env{
		GoogleClientId:      ReadEnvVar("GOOGLE_CLIENT_ID", ""),
		GoogleClientSecret:  ReadEnvVar("GOOGLE_CLIENT_SECRET", ""),
		AllowedRedirectUrls: strings.Split(ReadEnvVar("ALLOWED_REDIRECT_URLS", "http://localhost:3000/oauth/google/callback"), ","),
		Port:                ReadEnvVar("PORT", "3000"),
		DbHost:              ReadEnvVar("DB_HOST", "localhost"),
		DbPort:              ReadEnvVar("DB_PORT", "5432"),
		DbUser:              ReadEnvVar("DB_USER", "postgres"),
		DbPassword:          ReadEnvVar("DB_PASSWORD", "postgres"),
		DbName:              ReadEnvVar("DB_NAME", "timesheet-go"),
	}
}

var Vars *Env = NewEnv()
