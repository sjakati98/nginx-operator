package deployments

import (
    
        appsv1 "k8s.io/api/apps/v1"
    
        corev1 "k8s.io/api/core/v1"
    
        metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    
)

// NewDeploymentForCR ...
func NewDeploymentForCR(r *web.example.com/v1alpha1.Nginx) *v1.Deployment{
    var e *v1.Deployment
    elemYaml := `
    {
	"kind": "Deployment",
	"apiVersion": "apps/v1",
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
		"replicas": 1,
		"selector": {
			"matchLabels": {
				"app": "-nginx",
				"release": ""
			}
		},
		"template": {
			"metadata": {
				"creationTimestamp": null,
				"labels": {
					"app": "-nginx",
					"chart": "nginx-3.4.0",
					"heritage": "Tiller",
					"release": ""
				},
				"annotations": {
					"prometheus.io/port": "9113",
					"prometheus.io/scrape": "true"
				}
			},
			"spec": {
				"containers": [
					{
						"name": "-nginx",
						"image": "docker.io/bitnami/nginx:1.16.0-debian-9-r69",
						"ports": [
							{
								"name": "http",
								"containerPort": 8080
							}
						],
						"resources": {},
						"livenessProbe": {
							"httpGet": {
								"path": "/",
								"port": "http"
							},
							"initialDelaySeconds": 30,
							"timeoutSeconds": 5,
							"failureThreshold": 6
						},
						"readinessProbe": {
							"httpGet": {
								"path": "/",
								"port": "http"
							},
							"initialDelaySeconds": 5,
							"timeoutSeconds": 3,
							"periodSeconds": 5
						},
						"imagePullPolicy": "IfNotPresent"
					},
					{
						"name": "metrics",
						"image": "docker.io/nginx/nginx-prometheus-exporter:0.1.0",
						"command": [
							"/usr/bin/exporter",
							"-nginx.scrape-uri",
							"http://127.0.0.1:8080/status"
						],
						"ports": [
							{
								"name": "metrics",
								"containerPort": 9113
							}
						],
						"resources": {},
						"livenessProbe": {
							"httpGet": {
								"path": "/metrics",
								"port": "metrics"
							},
							"initialDelaySeconds": 15,
							"timeoutSeconds": 5
						},
						"readinessProbe": {
							"httpGet": {
								"path": "/metrics",
								"port": "metrics"
							},
							"initialDelaySeconds": 5,
							"timeoutSeconds": 1
						},
						"imagePullPolicy": "IfNotPresent"
					}
				]
			}
		},
		"strategy": {}
	},
	"status": {}
}
    `
    // Unmarshal Specified JSON to Kubernetes Resource
    err := json.Unmarshal([]byte(elemYaml), e)
    if err != nil {
        panic(err)
    }
    return e
}
