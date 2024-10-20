package test

import "testing"

func AssertEquals(expected any, actual any, t *testing.T) {
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
	}
}
