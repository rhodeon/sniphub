package main

import (
	"github.com/rhodeon/sniphub/pkg/testhelpers"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var mockHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("next handler called"))
})

func Test_secureHeaders(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	secureHeaders(mockHandler).ServeHTTP(rr, req)
	rs := rr.Result()

	t.Run("check headers", func(t *testing.T) {
		tests := []struct {
			key   string
			value string
		}{
			{"X-Frame-Options", "deny"},
			{"X-XSS-Protection", "1; mode=block"},
		}

		for _, tt := range tests {
			t.Run(tt.key, func(t *testing.T) {
				got := rs.Header.Get(tt.key)
				testhelpers.AssertString(t, got, tt.value)
			})
		}
	})

	t.Run("call next handler", func(t *testing.T) {
		if rs.StatusCode != http.StatusOK {
			t.Errorf("\nGot:\t%d\nWant:\t%d", rs.StatusCode, http.StatusOK)
		}

		rsBody, err := io.ReadAll(rs.Body)
		if err != nil {
			t.Fatal(err)
		}
		defer rs.Body.Close()
		gotBody := string(rsBody)
		wantBody := "next handler called"

		testhelpers.AssertString(t, gotBody, wantBody)
	})
}
