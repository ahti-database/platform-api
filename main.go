package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Jason-CKY/ahti/pkg/handlers"
	"github.com/Jason-CKY/ahti/pkg/schemas"
	"github.com/Jason-CKY/ahti/pkg/utils"
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func getKubernetesConfig() *rest.Config {
	var config *rest.Config
	home := homedir.HomeDir()
	kubeconfig := filepath.Join(home, ".kube", "config")
	if _, err := os.Stat(kubeconfig); err == nil {
		// if $HOME/.kube/config file exists
		log.Infof("Using kubeconfig file at %v", kubeconfig)
		// use the current context in kubeconfig
		_config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			panic(err.Error())
		}
		config = _config
	} else {
		log.Infof("Using incluster config...")
		_config, err := rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
		config = _config
	}
	return config
}

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

	// create the clientset
	config := getKubernetesConfig()
	utils.ClientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	e := echo.New()
	e.Validator = &schemas.APIValidator{Validator: validator.New()}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from the ahti platform API!")
	})
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "healthy")
	})

	g := e.Group("/api/v1")
	g.GET("/databases", handlers.ListDatabases)
	g.POST("/organizations/:organization/databases", handlers.CreateDatabase)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", utils.Port)))
}
