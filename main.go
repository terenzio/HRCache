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

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}

	username := os.Getenv("MARIADB_USER")
	if username == "" {
		log.Fatal("MARIADB_USER must be set")
	}

	password := os.Getenv("MARIADB_PASSWORD")
	if password == "" {
		log.Fatal("MARIADB_PASSWORD must be set")
	}

	host := os.Getenv("MARIADB_HOST")
	if host == "" {
		log.Fatal("MARIADB_HOST must be set")
	}

	database := os.Getenv("MARIADB_DATABASE")
	if database == "" {
		log.Fatal("MARIADB_DATABASE must be set")
	}

	fmt.Println("username:", username)
	fmt.Println("password:", password)
	fmt.Println("host:", host)
	fmt.Println("database:", database)

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, host, database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize the database adapter.
	dbAdapter := NewDatabaseAdapter(dsn)

	router := gin.Default()
	cache := NewCache()
	router.Use(CacheMiddleware(cache))
	router.POST("/employees", createEmployeeHandler(dbAdapter))
	router.GET("/employees/:id", getEmployeeHandler(dbAdapter))
	router.PUT("/employees/:id", updateEmployeeHandler(dbAdapter))
	router.DELETE("/employees/:id", deleteEmployeeHandler(dbAdapter))

	router.Run(":8080")

}
