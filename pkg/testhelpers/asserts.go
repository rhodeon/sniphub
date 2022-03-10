package testhelpers

import "testing"

func AssertString(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("\nGot:\t%q\nWant:\t%q", got, want)
	}
}

func AssertInt(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("\nGot:\t%d\nWant:\t%d", got, want)
	}
}
