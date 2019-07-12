package services

import (
    
        appsv1 "k8s.io/api/apps/v1"
    
        corev1 "k8s.io/api/core/v1"
    
        metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    
)

// NewServiceForCR ...
func NewServiceForCR(r *web.example.com/v1alpha1.Nginx) *v1.Service{
    var e *v1.Service
    elemYaml := `
    {
	"kind": "Service",
	"apiVersion": "v1",
	"metadata": {
		"name": "-nginx",
		"creationTimestamp": null,
		"labels": {
			"app": "-nginx",
			"chart": "nginx-3.4.0",
			"heritage": "Tiller",
			"release": ""
		}
	},
	"spec": {
		"ports": [
			{
				"name": "http",
				"port": 80,
				"targetPort": "http"
			}
		],
		"selector": {
			"app": "-nginx"
		},
		"type": "LoadBalancer",
		"externalTrafficPolicy": "Cluster"
	},
	"status": {
		"loadBalancer": {}
	}
}
    `
    // Unmarshal Specified JSON to Kubernetes Resource
    err := json.Unmarshal([]byte(elemYaml), e)
    if err != nil {
        panic(err)
    }
    return e
}
