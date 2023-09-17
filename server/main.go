package main

import (
	"github.com/joho/godotenv"
	"fmt"
	"os"
	"log"
	"host/core"
)

type environment struct {
	Url string
}

/**
*	Main function
*
*	@returns void
*/
func main() {
	env, err := loadEnv();
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	};

	fmt.Println("Server running on port: " + env.Url);

	core.HandleServerStart(env.Url);
}

/**
*	Load environment variables from .env file
*
*	@returns environment struct
*	@returns error
*/
func loadEnv() (environment, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file %v", err)
	};

	env := environment{
		Url: os.Getenv("SERVER_URL"),
	};

	return env, nil;
}
