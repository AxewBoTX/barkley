package api

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"

	"dondu/lib"
)

func Todo(router fiber.Router, DB *sql.DB) {
	router.Get("/todo", func(c *fiber.Ctx) error {
		rows, rows_fetch_err := DB.Query("SELECT * FROM Todos")
		if rows_fetch_err != nil {
			log.Fatal("Rows Fetch Error:", rows_fetch_err)
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		defer rows.Close()

		var Todos []lib.Todo
		for rows.Next() {
			var todo lib.Todo
			if row_scan_err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done); row_scan_err != nil {
				log.Fatal("Row Scan Error:", row_scan_err)
				return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}
			Todos = append(Todos, todo)
		}

		return c.JSON(Todos)
	})
	router.Post("/todo", func(c *fiber.Ctx) error {
		return c.JSON("Create Todo")
	})
	router.Patch("/todo/:todoID", func(c *fiber.Ctx) error {
		return c.JSON("Update Todo with ID")
	})
	router.Delete("/todo/:todoID", func(c *fiber.Ctx) error {
		return c.JSON("Delete Todo with ID")
	})
}
