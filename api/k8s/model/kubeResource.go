package model

import "k8s.io/apimachinery/pkg/runtime/schema"

type KubeResource struct {
	Client          schema.GroupVersionResource
	IsClusterScoped bool
}
