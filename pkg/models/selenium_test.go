package models

import (
	"testing"

	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/constants"
)

func TestCreatesSessionNegative(t *testing.T) {
	actual := CreateSession(nil, "http://127.0.0.01")
	if actual.Err == nil || actual.ResponseCode != constants.Unable_TO_CREATE_CLIENT_OBJECT {
		t.Fail()
	}
}

func TestDeleteSessionNegative(t *testing.T) {
	actual := DeleteSession("", "http://127.0.0.011")
	if actual.Err == nil || actual.ResponseCode != 422 {
		t.Fail()
	}
}
