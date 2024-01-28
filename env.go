package main

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv takes the path of a .env file and uses it to set environment variables
func LoadEnv(fileName string) {
	if err := godotenv.Load(fileName); err != nil {
		log.Fatal("No .env file found. Exiting", err)
	}
}
