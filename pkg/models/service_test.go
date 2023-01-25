package models

import "testing"

func TestIFDefaultValuesSetService(t *testing.T) {
	service := Service{}
	service.SetValues()
	if service.Kind != SERVICE_KIND || service.APIVersion != SERVICE_API_VERSION {
		t.Fail()
	}
}
