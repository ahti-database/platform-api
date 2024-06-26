package schemas

import (
	"context"

	"github.com/Jason-CKY/ahti/pkg/utils"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func int32Ptr(i int32) *int32 { return &i }

type Database struct {
	Id   string `json:"id"`
	Name string `json:"name" validate:"required"`
}

func (database Database) Create(namespace string) error {
	// TODO: create namespace if doesn't exist
	// TODO: define deployment
	deploymentsClient := utils.ClientSet.AppsV1().Deployments(namespace)
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
	// TODO: create deployment
	log.Info("Creating deployment...")
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	log.Infof("Created deployment %q.\n", result.GetObjectMeta().GetName())
	// TODO: create service
	// TODO: create ingress

	return nil
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
