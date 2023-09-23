package main

import (
	"fmt"
	"host/core"
	"host/database"
	"log"
	"os"

	"github.com/joho/godotenv"
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

	defer func() {
		if err := database.CloseDatabaseConnection(); err != nil {
			log.Fatalf("Error while closing database connection: %v", err)
		}
	}()
	
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
