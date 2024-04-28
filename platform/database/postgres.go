package database

import (
	"fmt"
	"new-go-api/app/models"
	"new-go-api/pkg/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/jackc/pgx/v4/stdlib" // load pgx driver for PostgreSQL
)

// PostgreSQLConnection func for connection to PostgreSQL database.
func PostgreSQLConnection() (*gorm.DB, error) {
	url, err := utils.ConnectionURLBuilder("postgres")

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	db.AutoMigrate(&models.User{}, &models.Post{})

	if err != nil {
		return nil, err
	}
	// Perform migrations only if the connection is successfully established
	// if err := db.AutoMigrate(&models.User{}); err != nil {
	// 	return nil, err
	// }
	if err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}

	fmt.Print("\nConnected to the database!")

	return db, nil

}
