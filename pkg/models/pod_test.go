package models

import "testing"

func TestIFDefaultValuesSet(t *testing.T) {
	pod := Pod{}
	pod.SetValues()
	if pod.Kind != POD_KIND || pod.APIVersion != POD_API_VERSION {
		t.Fail()
	}
	if len(pod.Spec.Containers) != 2 {
		t.Fail()
	}
}
