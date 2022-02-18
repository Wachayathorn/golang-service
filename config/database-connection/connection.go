package database_connection

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	models "github.com/wachayathorn/golang-service/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PG_DB *sql.DB
var GORM_DB *gorm.DB

func ConnectDatabaseByPQ() {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "admin")
	dbPassword := getEnv("DB_PASSWORD", "P@ssw0rd")
	dbName := getEnv("DB_NAME", "golang")

	connStr := fmt.Sprintf("host=%s  port=%s  user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic("Failed to connect to database with driver")
	}

	PG_DB = db
}

func ConnectDatabaseByGORM() {

	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "admin")
	dbPassword := getEnv("DB_PASSWORD", "P@ssw0rd")
	dbName := getEnv("DB_NAME", "golang")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database with gorm")
	}

	db.AutoMigrate(&models.User{})
	GORM_DB = db
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
