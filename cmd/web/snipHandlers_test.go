package main

import (
	"fmt"
	"github.com/rhodeon/sniphub/pkg/testhelpers"
	"net/http"
	"strings"
	"testing"
)

func Test_application_showSnip(t *testing.T) {
	app := newTestApp(t)
	testServer := newTestServer(t, app.routesHandler())
	defer testServer.Close()

	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{"valid id", "/snip/1", 200, "this is a mock snip"},
		{"zero id", "/snip/0", 404, fmt.Sprintln(http.StatusText(http.StatusNotFound))},
		{"id out of range", "/snip/10", 404, fmt.Sprintln(http.StatusText(http.StatusNotFound))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// assert status code
			code, _, body := testServer.get(t, tt.urlPath)
			testhelpers.AssertInt(t, code, tt.wantCode)

			// assert response body
			if !strings.Contains(body, tt.wantBody) {
				t.Errorf("want body to contain %q", tt.wantBody)
			}
		})
	}
}
