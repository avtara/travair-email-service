package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type EmailEnv struct {
	ConfigSmtpHost     string
	ConfigSmtpPort     int
	ConfigSenderName   string
	ConfigAuthEmail    string
	ConfigAuthPassword string
}

func GetEmailEnv() *EmailEnv {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	mailPort, _ := strconv.Atoi(os.Getenv("CONFIG_SMTP_PORT"))
	return &EmailEnv{
		ConfigSmtpHost:     os.Getenv("CONFIG_SMTP_HOST"),
		ConfigSmtpPort:     mailPort,
		ConfigSenderName:   os.Getenv("CONFIG_SENDER_NAME"),
		ConfigAuthEmail:    os.Getenv("CONFIG_AUTH_EMAIL"),
		ConfigAuthPassword: os.Getenv("CONFIG_AUTH_PASSWORD"),
	}
}
