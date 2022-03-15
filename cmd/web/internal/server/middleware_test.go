package server

import (
	"context"
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
	req, err := http.NewRequest(http.MethodGet, homeRoute, nil)
	testhelpers.AssertFatalError(t, err)

	secureHeaders(mockHandler).ServeHTTP(rr, req)
	rs := rr.Result()

	t.Run("check headers", func(t *testing.T) {
		tests := []testHeader{
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
		testhelpers.AssertInt(t, rs.StatusCode, http.StatusOK)

		rsBody, err := io.ReadAll(rs.Body)
		testhelpers.AssertFatalError(t, err)
		defer rs.Body.Close()

		gotBody := string(rsBody)
		wantBody := "next handler called"
		testhelpers.AssertString(t, gotBody, wantBody)
	})
}

func TestApplication_requireAuthentication(t *testing.T) {
	app := newTestApp(t)

	tests := []struct {
		name            string
		isAuthenticated bool
		wantCode        int
		wantHeaders     []testHeader
	}{
		{"authenticated user", true, http.StatusOK, []testHeader{{"Cache-Control", "no-store"}}},
		{"unauthenticated user", false, http.StatusSeeOther, []testHeader{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, homeRoute, nil)
			testhelpers.AssertFatalError(t, err)

			if tt.isAuthenticated {
				// emulate logged in user
				ctx := context.WithValue(req.Context(), ContextKeyIsAuthenticated, true)
				req = req.WithContext(ctx)
			}

			// loadAndSave initializes the session needed by requireAuthentication
			app.SessionManager.LoadAndSave(app.requireAuthentication(mockHandler)).ServeHTTP(rr, req)
			rs := rr.Result()

			// assert status code
			code := rs.StatusCode
			testhelpers.AssertInt(t, code, tt.wantCode)

			// assert headers
			for _, header := range tt.wantHeaders {
				t.Run(header.key+" header", func(t *testing.T) {
					got := rs.Header.Get(header.key)
					testhelpers.AssertString(t, got, header.value)
				})
			}
		})
	}
}
