package api

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func Todo(router fiber.Router, DB *sql.DB) {
	router.Get("/todo", func(c *fiber.Ctx) error {
		return c.JSON("Get Todo List")
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
