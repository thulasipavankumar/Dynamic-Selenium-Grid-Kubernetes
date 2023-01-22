package models

import (
	"testing"
)

func TestSomething(t *testing.T) {
	actual := CreateSession(nil, "http://127.0.0.011")
	if actual.err == nil || actual.responseCode != 422 {
		t.Fail()
	}
}
