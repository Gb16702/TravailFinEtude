package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB;

/**
*	This function connects to the database
*
*	@param string connectionString
*	@returns void
*/
func ConnectToDatabase(connectionString string) {
	var err error;
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	});

	if err != nil {
		log.Fatalf("Error while connecting to database: %v", err)
	}
	
	log.Println("Successfully connected to database");
}

/**
*	This function migrates the schemas
*
*	@param *gorm.DB DB
*	@returns error
*/
func migrateSchemas(DB *gorm.DB) error {

	return nil;
}
