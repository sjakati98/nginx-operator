package apis

import (
	"github.com/redhat-nfvpe/helm2go-operator-sdk/nginx-operator/pkg/apis/web/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v1alpha1.SchemeBuilder.AddToScheme)
}
