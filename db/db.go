package db

import (
	"block_chain/system"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"log"
)

var db *gorm.DB
var err error

// InitDB ...
func InitDB() (*gorm.DB, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		system.GetConfiguration().PgUser,
		system.GetConfiguration().PgPassword,
		system.GetConfiguration().PgHost,
		system.GetConfiguration().PgPort,
		system.GetConfiguration().PgDatabase,
	)

	db, err = gorm.Open("postgres", dbinfo)

	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}

	log.Println("Database connected")

	return db, err
}

// GetDB ...
func GetDB() *gorm.DB {
	return db
}
