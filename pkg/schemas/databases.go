package schemas

import (
	"context"

	"github.com/Jason-CKY/ahti/pkg/utils"
	"github.com/google/uuid"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Database struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func ListDatabases() ([]Database, error) {
	allConfigmaps, err := utils.ClientSet.CoreV1().ConfigMaps("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var databases []Database

	for _, cm := range allConfigmaps.Items {
		databases = append(databases, Database{
			Id:   uuid.New().String(),
			Name: cm.Name,
		})
	}

	return databases, nil
}
