package handlers

import (
	"net/http"

	"github.com/Jason-CKY/ahti/pkg/schemas"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func ListDatabases(c echo.Context) error {
	database := schemas.Database{
		Id:   uuid.New().String(),
		Name: "test database",
	}
	return c.JSON(http.StatusOK, database)
}
