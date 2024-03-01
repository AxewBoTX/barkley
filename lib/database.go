package lib

import (
	"database/sql"
	"os"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

func PrepareDatabase() *sql.DB {
	if _, folder_check_err := os.Stat("database"); os.IsNotExist(folder_check_err) {
		folder_create_err := os.Mkdir("database", 0755)
		if folder_create_err != nil {
			log.Fatal("Failed To Create Folder", "Error", folder_create_err)
		}
	}
	if _, file_check_err := os.Stat(DB_PATH); os.IsNotExist(file_check_err) {
		_, file_create_err := os.Create(DB_PATH)
		if file_create_err != nil {
			log.Fatal("Failed To Create DB File", "Error", file_create_err)
		}
	}
	DB, db_open_err := sql.Open("sqlite", DB_PATH)
	if db_open_err != nil {
		log.Fatal("Failed To Open Database", "Error", db_open_err)
	}
	return DB
}

func HandleMigrations(DB *sql.DB) {
	if _, table_create_err := DB.Exec(`CREATE TABLE IF NOT EXISTS Todos(
		id TEXT PRIMARY KEY,
		title TEXT
	);`); table_create_err != nil {
		log.Fatal("Failed To Create Table", "Error", table_create_err)
	}
}

func GenerateRandomRows(DB *sql.DB) {
	if _, row_create_err := DB.Exec(
		`INSERT INTO Todos (id,title) VAlUES (?,?)`,
		uuid.NewString(),
		"Go Shopping",
	); row_create_err != nil {
		log.Fatal("Failed To Create Row", "Error", row_create_err)
	}
	if _, row_create_err := DB.Exec(
		`INSERT INTO Todos (id,title) VALUES (?,?)`,
		uuid.NewString(),
		"Bath yourself",
	); row_create_err != nil {
		log.Fatal("Failed To Create Row", "Error", row_create_err)
	}
}
