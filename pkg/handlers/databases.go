package handlers

import (
	"net/http"

	"github.com/Jason-CKY/ahti/pkg/schemas"
	"github.com/labstack/echo/v4"
)

func ListDatabases(c echo.Context) error {
	databases, err := schemas.ListDatabases()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, schemas.HTTPErr{Detail: err.Error()})
	}
	return c.JSON(http.StatusOK, databases)
}
