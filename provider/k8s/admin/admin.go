package admin

import (
	"k8s.io/client-go/kubernetes"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func NewAdmin(cs *kubernetes.Clientset) *Admin {
	return &Admin{
		corev1: cs.CoreV1(),
	}
}

type Admin struct {
	corev1 corev1.CoreV1Interface
}
