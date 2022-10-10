package lib

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(env Env) Database {
	url := env.DB_URL

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	log.Println("Database connection established")

	return Database{
		DB: db,
	}
}
