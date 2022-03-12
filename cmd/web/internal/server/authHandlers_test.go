package server

import (
	"github.com/rhodeon/sniphub/pkg/forms"
	"github.com/rhodeon/sniphub/pkg/testhelpers"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestApplication_signupUserPost(t *testing.T) {
	app := newTestApp(t)
	testServer := newTestServer(t, app.RouteHandler())
	defer testServer.Close()

	_, _, body := testServer.get(t, signupRoute)
	csrfToken := extractCSRFToken(t, []byte(body))

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
		{"empty username", "", "rhodeon@mail.com", "passworder", csrfToken, http.StatusOK, forms.ErrBlankField},
		{"empty email", "rhodeon", "", "passworder", csrfToken, http.StatusOK, forms.ErrBlankField},
		{"empty password", "rhodeon", "rhodeon@mail.com", "", csrfToken, http.StatusOK, forms.ErrBlankField},
		{"mismatched email (missing @)", "rhodeon", "rhodeonmail.com", "passworder", csrfToken, http.StatusOK, forms.ErrInvalidField},
		{"mismatched email (missing local name)", "rhodeon", "@mail.com", "passworder", csrfToken, http.StatusOK, forms.ErrInvalidField},
		{"mismatched email (missing period prefix in domain)", "rhodeon", "rhodeon@.com", "passworder", csrfToken, http.StatusOK, forms.ErrInvalidField},
		{"mismatched email (missing period suffix in domain)", "rhodeon", "rhodeon@mail.", "passworder", csrfToken, http.StatusOK, forms.ErrInvalidField},
		{"username over max length", "ddsssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssddsssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssddsssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss", "rhodeon@mail.com", "passworder", csrfToken, http.StatusOK, "This field must not have over 255 characters"},
		{"email over max length", "rhodeon", "ddsssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssddsssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssddsssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss", "passworder", csrfToken, http.StatusOK, "This field must not have over 255 characters"},
		{"password below minimum length", "rhodeon", "rhodeon@mail.com", "pass", csrfToken, http.StatusOK, "This field must have at least 10 characters"},
		{"username already exists", "rhodeon", "ruona@mail.com", "passworder", csrfToken, http.StatusOK, forms.ErrExistingUsername},
		{"email already exists", "ruona", "rhodeon@mail.com", "passworder", csrfToken, http.StatusOK, forms.ErrExistingEmail},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{
				forms.Username:  {tt.userName},
				forms.Email:     {tt.userEmail},
				forms.Password:  {tt.userPassword},
				forms.CsrfToken: {tt.csrfToken},
			}

			code, _, body := testServer.postForm(t, signupRoute, form)

			// assert response status code
			testhelpers.AssertInt(t, code, tt.wantCode)

			// assert response body
			if !strings.Contains(body, tt.wantBody) {
				t.Errorf("want body to contain %q", tt.wantBody)
			}
		})
	}
}

func TestApplication_loginUserPost(t *testing.T) {
	app := newTestApp(t)
	testServer := newTestServer(t, app.RouteHandler())
	defer testServer.Close()

	_, _, body := testServer.get(t, loginRoute)
	csrfToken := extractCSRFToken(t, []byte(body))
	t.Logf("Token: %s", csrfToken)

	tests := []struct {
		name      string
		email     string
		password  string
		csrfToken string
		wantCode  int
		wantBody  string
	}{
		{"valid credentials", "rhodeon@mail.com", "qwerty123456", csrfToken, http.StatusSeeOther, ""},
		{"wrong email", "rhodeos@mail.com", "qwerty123456", csrfToken, http.StatusOK, forms.ErrInvalidEmailOrPassword},
		{"wrong password", "rhodeon@mail.com", "zzz", csrfToken, http.StatusOK, forms.ErrInvalidEmailOrPassword},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{
				forms.Email:     {tt.email},
				forms.Password:  {tt.password},
				forms.CsrfToken: {tt.csrfToken},
			}

			code, _, body := testServer.postForm(t, loginRoute, form)

			// assert status code
			testhelpers.AssertInt(t, code, tt.wantCode)

			// assert response body
			if !strings.Contains(body, tt.wantBody) {
				t.Errorf("want body to contain %q", tt.wantBody)
			}
		})
	}
}
