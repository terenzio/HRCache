package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	// Load the environment variables from the .env file
	// and check that the required variables are set.
	if err := loadEnv(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Setup the database connection
	// and check that the connection is successful.
	dsn, err := setupDatabase()
	if err != nil {
		log.Fatalf("Failed to setup database DSN: %v", err)
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create a new DatabaseAdapter with the database connection.
	dbAdapter := NewDatabaseAdapter(dsn)

	// Setup the Gin router and start the server.
	// The router is configured with the database adapter
	// to handle the employee CRUD operations.
	// The server listens on port 8080.
	router := setupRouter(dbAdapter)
	router.Run(":8080")

}

func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	envVars := []string{"MARIADB_USER", "MARIADB_PASSWORD", "MARIADB_HOST", "MARIADB_DATABASE"}
	for _, envVar := range envVars {
		if os.Getenv(envVar) == "" {
			return fmt.Errorf("%s must be set", envVar)
		}
	}
	return nil
}

func setupDatabase() (string, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		os.Getenv("MARIADB_USER"),
		os.Getenv("MARIADB_PASSWORD"),
		os.Getenv("MARIADB_HOST"),
		os.Getenv("MARIADB_DATABASE"),
	)
	return dsn, nil
}

func setupRouter(dbAdapter *DatabaseAdapter) *gin.Engine {
	router := gin.Default()
	cache := NewCache()

	// Use the cache middleware for all routes
	// to cache the responses for 30 seconds.
	router.Use(CacheMiddleware(cache))

	router.POST("/employees", createEmployeeHandler(dbAdapter))
	router.GET("/employees/:id", getEmployeeHandler(dbAdapter))
	router.PUT("/employees/:id", updateEmployeeHandler(dbAdapter))
	router.DELETE("/employees/:id", deleteEmployeeHandler(dbAdapter))
	return router
}
