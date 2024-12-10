package db

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var dbClient *gorm.DB = nil

func Connect(dbUri string) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dbUri,
		DefaultStringSize:         256,
		SkipInitializeWithVersion: false,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
	}))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to the database!")
	dbClient = db
	return dbClient
}
