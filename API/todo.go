package api

import (
	"database/sql"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"dondu/lib"
)

func Todo(router fiber.Router, DB *sql.DB) {
	// Get Todo List
	router.Get("/todo", func(c *fiber.Ctx) error {
		rows, rows_fetch_err := DB.Query("SELECT * FROM Todos")
		if rows_fetch_err != nil {
			log.Println("Rows Fetch Error:", rows_fetch_err)
			return c.Status(fiber.StatusInternalServerError).SendString("Database Row Fetch Error")
		}
		defer rows.Close()

		var Todos []lib.Todo
		for rows.Next() {
			var todo lib.Todo
			if row_scan_err := rows.Scan(&todo.ID, &todo.Title); row_scan_err != nil {
				log.Println("Row Scan Error:", row_scan_err)
				return c.Status(fiber.StatusInternalServerError).
					SendString("Database Row Scan Error")
			}
			Todos = append(Todos, todo)
		}

		return c.Status(fiber.StatusOK).JSON(Todos)
	})

	// Create Todo
	router.Post("/todo", func(c *fiber.Ctx) error {
		var todo lib.Todo
		if body_parse_err := c.BodyParser(&todo); body_parse_err != nil {
			log.Println("Body Parse Error:", body_parse_err)
			return c.Status(fiber.StatusInternalServerError).SendString("Request Body Parse Error")
		}
		todo.ID = uuid.NewString()
		if _, row_create_err := DB.Exec(`INSERT INTO Todos (id,title) VALUES(?,?)`, todo.ID, todo.Title); row_create_err != nil {
			log.Println("Row Create Error:", row_create_err)
			return c.Status(fiber.StatusInternalServerError).SendString("Database Row Create Error")
		}
		return c.Status(fiber.StatusOK).JSON(todo)
	})

	// Delete Todo
	router.Delete("/todo/:todoID", func(c *fiber.Ctx) error {
		todoID := c.Params("todoID")
		if len(todoID) == 0 || len(strings.TrimSpace(todoID)) == 0 ||
			strings.TrimSpace(todoID) == "" {
			return c.Status(fiber.StatusBadRequest).JSON("Bad Request ID")
		} else {
			if _, row_delete_err := DB.Exec(`DELETE FROM Todos WHERE id = ?`, todoID); row_delete_err != nil {
				log.Println("Row Delete Error:", row_delete_err)
				return c.Status(fiber.StatusInternalServerError).
					SendString("Database Row Delete Error")
			} else {
				return c.Status(fiber.StatusOK).JSON("SUCCESS")
			}
		}
	})
}
