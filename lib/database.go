package lib

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/pocketbase/dbx"
	_ "modernc.org/sqlite"
)

// handle database file and object creation
func CreateDatabase() *dbx.DB {
	db_file_path := filepath.Join(DB_FOLDER, DB_FILE)

	// check if the DB_FOLDER exists and if not then create it
	if _, folder_check_err := os.Stat(DB_FOLDER); os.IsNotExist(folder_check_err) {
		log.Warnf("Folder %s does not exist, creating it...", DB_FOLDER)
		folder_create_err := os.Mkdir(DB_FOLDER, 0755)
		if folder_create_err != nil {
			log.Fatal("Failed to create folder", "Error", folder_create_err)
		}
	}

	// check if the DB_FILE exists and if not then create it
	if _, file_check_err := os.Stat(db_file_path); os.IsNotExist(file_check_err) {
		log.Warnf("File %s does not exist, creating it...", db_file_path)
		file, file_create_err := os.Create(db_file_path)
		if file_create_err != nil {
			log.Fatal("Failed to create file", "Error", file_create_err)
		}
		file.Close()
	}

	// create the dbx.DB object using SQLite driver
	db, database_create_err := dbx.Open("sqlite", db_file_path)
	if database_create_err != nil {
		log.Fatal("Failed to create database", "Error", database_create_err)
	}

	return db
}

func HandleMigrations(DB *dbx.DB) {
	// TODO: create tables in the database file
}
