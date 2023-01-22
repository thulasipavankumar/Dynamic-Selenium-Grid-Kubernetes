package models

import (
	"testing"
)

func TestNegativeDeployment(t *testing.T) {
	deployment := Deployment{}
	details := deployment.GetDetails()
	if details.IngressName != "" || details.PodName != "" || details.ServiceName != "" {
		t.Fail()
	}
}
