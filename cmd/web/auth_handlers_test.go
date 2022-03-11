package main

import (
	"github.com/rhodeon/sniphub/pkg/forms"
	"github.com/rhodeon/sniphub/pkg/testhelpers"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func Test_application_signupUserPost(t *testing.T) {
	app := newTestApp(t)
	testServer := newTestServer(t, app.routesHandler())
	defer testServer.Close()

	_, _, _ = testServer.get(t, "/auth/signup")
	csrfToken := "" // extractCSRFToken(t, []byte(body))

	tests := []struct {
		name         string
		userName     string
		userEmail    string
		userPassword string
		csrfToken    string
		wantCode     int
		wantBody     string
	}{
		{"valid details", "ruona", "ruona@mail.com", "passworder", csrfToken, http.StatusSeeOther, ""},
		{"empty username", "", "rhodeon@mail.com", "passworder", csrfToken, http.StatusOK, "This field cannot be blank"},
		{"empty email", "rhodeon", "", "passworder", csrfToken, http.StatusOK, "This field cannot be blank"},
		{"empty password", "rhodeon", "rhodeon@mail.com", "", csrfToken, http.StatusOK, "This field cannot be blank"},
		{"mismatched email (missing @)", "rhodeon", "rhodeonmail.com", "passworder", csrfToken, http.StatusOK, "This field is invalid"},
		{"mismatched email (missing local name)", "rhodeon", "@mail.com", "passworder", csrfToken, http.StatusOK, "This field is invalid"},
		{"mismatched email (missing period prefix in domain)", "rhodeon", "rhodeon@.com", "passworder", csrfToken, http.StatusOK, "This field is invalid"},
		{"mismatched email (missing period suffix in domain)", "rhodeon", "rhodeon@mail.", "passworder", csrfToken, http.StatusOK, "This field is invalid"},
		{"username over max length", "ddsssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssddsssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssddsssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss", "rhodeon@mail.com", "passworder", csrfToken, http.StatusOK, "This field must not have over 255 characters"},
		{"email over max length", "rhodeon", "ddsssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssddsssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssddsssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss", "passworder", csrfToken, http.StatusOK, "This field must not have over 255 characters"},
		{"password below minimum length", "rhodeon", "rhodeon@mail.com", "pass", csrfToken, http.StatusOK, "This field must have at least 10 characters"},
		{"username already exists", "rhodeon", "ruona@mail.com", "passworder", csrfToken, http.StatusOK, "Username is already taken"},
		{"email already exists", "ruona", "rhodeon@mail.com", "passworder", csrfToken, http.StatusOK, "Email already in use"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{
				forms.Username:  {tt.userName},
				forms.Email:     {tt.userEmail},
				forms.Password:  {tt.userPassword},
				forms.CsrfToken: {tt.csrfToken},
			}

			code, _, body := testServer.postForm(t, "/auth/signup", form)

			// assert response status code
			testhelpers.AssertInt(t, code, tt.wantCode)

			// assert response body
			if !strings.Contains(body, tt.wantBody) {
				t.Errorf("want body to contain %q", tt.wantBody)
			}
		})
	}
	// TODO: fix csrf token mismatch
}
