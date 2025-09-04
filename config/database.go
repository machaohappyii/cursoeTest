package config

import (
	"gin-user-api/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitDB(config *Config) (*gorm.DB, error) {
	db, err := gorm.Open(config.Database.Driver, config.GetDatabaseDSN())
	if err != nil {
		return nil, err
	}
	
	// Enable logging in debug mode
	if config.Server.Mode == "debug" {
		db = db.Debug()
	}
	
	// Configure connection pool for MySQL
	if config.Database.Driver == "mysql" {
		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(100)
	}
	
	// Auto-migrate the User model
	if err := db.AutoMigrate(&models.User{}).Error; err != nil {
		return nil, err
	}
	
	return db, nil
}