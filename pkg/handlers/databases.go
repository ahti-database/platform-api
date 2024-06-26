package handlers

import (
	"context"
	"net/http"

	"github.com/Jason-CKY/ahti/pkg/schemas"
	"github.com/Jason-CKY/ahti/pkg/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListDatabases(c echo.Context) error {

	allConfigmaps, err := utils.ClientSet.CoreV1().ConfigMaps("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	var databases []schemas.Database

	for _, cm := range allConfigmaps.Items {
		databases = append(databases, schemas.Database{
			Id:   uuid.New().String(),
			Name: cm.Name,
		})
	}

	return c.JSON(http.StatusOK, databases)
}
