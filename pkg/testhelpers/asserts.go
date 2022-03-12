// Package testhelpers provides functions for assisting with tests.
package testhelpers

import (
	"reflect"
	"strings"
	"testing"
)

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

func AssertError(t *testing.T, got error, want error) {
	t.Helper()

	if got != want {
		t.Errorf("\nGot Error:\t%+v\nWant Error:\t%+v", got, want)
	}
}

func AssertFatalError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatal(err)
	}
}

func AssertStruct(t *testing.T, got interface{}, want interface{}) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\nGot:\t%+v\nWant:\t%+v", got, want)
	}
}

// AssertTemplateContent verifies if a wanted string is contained in a template body.
func AssertTemplateContent(t *testing.T, body string, want string) {
	t.Helper()

	if !strings.Contains(body, want) {
		t.Errorf("want body to contain %q", want)
	}
}
