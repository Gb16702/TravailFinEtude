package main

import (
	"github.com/joho/godotenv"
	"fmt"
	"os"
	"log"
	"host/core"
	"host/database"
)

type environment struct {
	ServerUrl string
	ConnectionString string
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

	fmt.Println("Server running on port: " + env.ServerUrl);

	database.ConnectToDatabase(env.ConnectionString);
	core.HandleServerStart(env.ServerUrl);
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
		ServerUrl: os.Getenv("SERVER_URL"),
		ConnectionString: os.Getenv("DB_URL"),
	};

	return env, nil;
}
