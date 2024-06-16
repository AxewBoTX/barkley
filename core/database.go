package core

import (
	"context"

	"github.com/charmbracelet/log"

	"github.com/axewbotx/barkley/db"
)

type Database struct {
	Client  *db.PrismaClient
	Context context.Context
}

func NewDatabase() *Database {
	return &Database{}
}

func (DB *Database) Connect() *Database {
	client := db.NewClient()
	if client_connect_err := client.Prisma.Connect(); client_connect_err != nil {
		log.Fatal("Failed to connect database client", "Error", client_connect_err)
	}
	DB.Client = client
	DB.Context = context.Background()
	return DB
}
