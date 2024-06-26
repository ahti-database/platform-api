package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/Jason-CKY/ahti/pkg/handlers"
	"github.com/Jason-CKY/ahti/pkg/utils"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()

	if err != nil {
		log.Infof("Error loading .env file: %v\nUsing environment variables instead...", err)
	}

	flag.StringVar(&utils.LogLevel, "log-level", utils.LookupEnvOrString("LOG_LEVEL", utils.LogLevel), "Logging level for the server")
	flag.StringVar(&utils.Port, "port", utils.LookupEnvOrString("PORT", utils.Port), "Port number to serve API")

	flag.Parse()

	// setup logrus
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:          true,
		DisableLevelTruncation: true,
	})

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from the ahti platform API!")
	})
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "healthy")
	})

	g := e.Group("/api/v1")
	g.GET("/databases", handlers.ListDatabases)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", utils.Port)))
}
