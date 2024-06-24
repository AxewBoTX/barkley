package core

import (
	"github.com/charmbracelet/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	Client *gorm.DB
}

func NewDatabase(gorm_config gorm.Config) *Database {
	db_client, db_client_connect_err := gorm.Open(sqlite.Open(DB_FILE), &gorm_config)
	if db_client_connect_err != nil {
		log.Fatal("Failed to connect database client", "Error", db_client_connect_err)
	}
	return &Database{
		Client: db_client,
	}
}

func (DB *Database) HandleMigrations() {
	DB.Client.AutoMigrate(Project{}, Section{}, Task{}, SubTask{})
}
