package models

import (
	"testing"
)

func TestCreatesSessionNegative(t *testing.T) {
	actual := CreateSession(nil, "http://127.0.0.011")
	if actual.Err == nil || actual.ResponseCode != 422 {
		t.Fail()
	}
}

func TestDeleteSessionNegative(t *testing.T) {
	actual := DeleteSession("", "http://127.0.0.011")
	if actual.Err == nil || actual.ResponseCode != 422 {
		t.Fail()
	}
}
