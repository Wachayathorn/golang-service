package database_connection

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	models "github.com/wachayathorn/golang-service/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/go-pg/pg/v10"
)

var PG_DB *pg.DB
var GORM_DB *gorm.DB
var SQLX *sqlx.DB

func ConnectDatabaseByPQ() {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "admin")
	dbPassword := getEnv("DB_PASSWORD", "P@ssw0rd")
	dbName := getEnv("DB_NAME", "golang")

	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", dbHost, dbPort),
		User:     dbUser,
		Password: dbPassword,
		Database: dbName,
	})

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

func ConnectDatabaseBySQLX() {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "admin")
	dbPassword := getEnv("DB_PASSWORD", "P@ssw0rd")
	dbName := getEnv("DB_NAME", "golang")

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		panic("Failed to connect to database with sqlx")
	}

	SQLX = db
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
