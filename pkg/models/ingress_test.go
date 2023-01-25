package models

import "testing"

func TestIFDefaultValuesSetIngress(t *testing.T) {
	ingress := Ingress{}
	ingress.SetValues()
	if ingress.Kind != INGRESS_KIND || ingress.APIVersion != INGRESS_API_VERSION {
		t.Fail()
	}
}
