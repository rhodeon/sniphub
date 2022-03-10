package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_showSnip(t *testing.T) {
	t.Run("non-integer snip id generates 404 response", func(t *testing.T) {
		rr := httptest.NewRecorder()
		r, err := http.NewRequest(http.MethodGet, "/snip/first", nil)
		if err != nil {
			t.Fatal(err)
		}

		app := &application{}
		app.showSnip(rr, r)
		rs := rr.Result()

		// assert status code
		gotStatusCode := rs.StatusCode
		wantStatusCode := http.StatusNotFound

		if gotStatusCode != wantStatusCode {
			t.Errorf("\nGot:\t%d\nWant:\t%d", gotStatusCode, wantStatusCode)
		}

		// assert response body
		rsBody, err := io.ReadAll(rs.Body)
		if err != nil {
			t.Fatal(err)
		}
		defer rs.Body.Close()
		gotBody := string(rsBody)
		wantBody := http.StatusText(http.StatusNotFound) + "\n"

		if gotBody != wantBody {
			t.Errorf("\nGot:\t%q\nWant:\t%q", gotBody, wantBody)
		}
	})

	// TODO: Add subtests for database actions.
}
