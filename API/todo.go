package api

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"dondu/lib"
)

func Todo(router *echo.Group, DB *sql.DB) {
	// Get Todo List
	router.GET("/todo", func(c echo.Context) error {
		rows, rows_fetch_err := DB.Query("SELECT * FROM Todos")
		if rows_fetch_err != nil {
			log.Error("Failed To Fetch Rows", "Error", rows_fetch_err)
			return c.String(http.StatusInternalServerError, "Database Row Fetch Error")
		}
		defer func() {
			rows.Close()
		}()

		var Todos []lib.Todo
		for rows.Next() {
			var todo lib.Todo
			if row_scan_err := rows.Scan(&todo.ID, &todo.Title); row_scan_err != nil {
				log.Error("Failed To Scan Row", "Error", row_scan_err)
				return c.String(http.StatusInternalServerError, "Database Row Scan Error")
			}
			Todos = append(Todos, todo)
		}

		return c.JSON(http.StatusOK, Todos)
	})
	// Get Todo List
	router.GET("/todo", func(c echo.Context) error {
		rows, rows_fetch_err := DB.Query("SELECT * FROM Todos")
		if rows_fetch_err != nil {
			log.Error("Failed To Fetch Rows", "Error", rows_fetch_err)
			return c.String(http.StatusInternalServerError, "Database Row Fetch Error")
		}
		defer func() {
			rows.Close()
		}()

		var Todos []lib.Todo
		for rows.Next() {
			var todo lib.Todo
			if row_scan_err := rows.Scan(&todo.ID, &todo.Title); row_scan_err != nil {
				log.Error("Failed To Scan Row", "Error", row_scan_err)
				return c.String(http.StatusInternalServerError, "Database Row Scan Error")
			}
			Todos = append(Todos, todo)
		}

		return c.JSON(http.StatusOK, Todos)
	})

	// Create Todo
	router.POST("/todo", func(c echo.Context) error {
		var todo lib.Todo
		if todo_bind_err := c.Bind(&todo); todo_bind_err != nil {
		}
		todo.ID = uuid.NewString()
		if _, row_create_err := DB.Exec(`INSERT INTO Todos (id,title) VALUES(?,?)`, todo.ID, todo.Title); row_create_err != nil {
			log.Error("Failed To Fetch Row", "Error", row_create_err)
			return c.String(http.StatusInternalServerError, "Database Row Create Error")
		}
		return c.JSON(http.StatusOK, todo)
	})

	// Delete Todo
	router.DELETE("/todo/:todoID", func(c echo.Context) error {
		todoID := c.Param("todoID")
		if len(todoID) == 0 || len(strings.TrimSpace(todoID)) == 0 ||
			strings.TrimSpace(todoID) == "" {
			return c.JSON(http.StatusBadRequest, "Bad Request ID")
		} else {
			if _, row_delete_err := DB.Exec(`DELETE FROM Todos WHERE id = ?`, todoID); row_delete_err != nil {
				log.Error("Failed To Delete Row", "Error", row_delete_err)
				return c.JSON(http.StatusInternalServerError, "Database Row Delete Error")
			} else {
				return c.JSON(http.StatusOK, "SUCCESS")
			}
		}
	})
}
