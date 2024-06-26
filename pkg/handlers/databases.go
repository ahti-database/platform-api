package handlers

import (
	"net/http"

	"github.com/Jason-CKY/ahti/pkg/schemas"
	"github.com/labstack/echo/v4"
)

func ListDatabases(c echo.Context) error {
	databases, err := schemas.ListDatabases()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, databases)
}

func CreateDatabase(c echo.Context) error {
	namespace := c.Param("organization")
	var database schemas.Database
	// TODO: validate https://echo.labstack.com/docs/request#validate-data
	if err := c.Bind(&database); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(database); err != nil {
		return err
	}
	if err := database.Create(namespace); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, database)
}
