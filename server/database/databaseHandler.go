package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"host/database/models"
)

var DB *gorm.DB;

/**
*	This function connects to the database
*
*	@param string connectionString
*	@returns void
*/
func ConnectToDatabase(connectionString string) {
	defer func() {
		if err := CloseDatabaseConnection(DB); err != nil {
			log.Fatalf("Error while closing database connection: %v", err)
		}
	}()

	var err error;
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	});

	if err != nil {
		log.Fatalf("Error while connecting to database: %v", err)
	}

	log.Println("Successfully connected to database");

	message, err := migrateSchemas(DB);
	if err != nil {
		log.Fatalf("Error while migrating schemas: %v", err)
	}

	log.Println(message);
}

/**
*	This function migrates the schemas
*
*	@param *gorm.DB database
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
*	@param *gorm.DB database
*	@returns error
*/
func CloseDatabaseConnection(database *gorm.DB) error  {
	db, err := database.DB();
	if err != nil {
		return err;
	}

	db.Close();

	return nil;
}
