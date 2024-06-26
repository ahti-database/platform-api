package utils

import "k8s.io/client-go/kubernetes"

var (
	LogLevel  = "info"
	Port      = "8080"
	ClientSet *kubernetes.Clientset
)
