package database

import (
	"log"

	"host/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB;

/**
*	This function connects to the database
*
*	@param connectionString string
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

	log.Println("Successfully connected to database", DB.Name());

	message, err := migrateSchemas(DB);
	if err != nil {
		log.Fatalf("Error while migrating schemas: %v", err)
	}

	log.Println(message);
}

/**
*	This function migrates the schemas
*
*	@param database *gorm.DB
*	@returns string
*	@returns error
*/
func migrateSchemas(database *gorm.DB) (string, error) {
	if err := database.AutoMigrate(&models.User{}); err != nil {
		return "", err;
	}
	return "Successfully migrated schemas", nil;
}

/**
*	This function closes the database connection
*
*	@returns error
*/
func CloseDatabaseConnection() error  {
	db, err := DB.DB();
	if err != nil {
		return err;
	}
	db.Close();

	return nil;
}
