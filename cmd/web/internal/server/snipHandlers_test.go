package server

import (
	"github.com/rhodeon/sniphub/pkg/forms"
	"github.com/rhodeon/sniphub/pkg/testhelpers"
	"net/http"
	"net/url"
	"testing"
)

func TestApplication_showSnip(t *testing.T) {
	app := newTestApp(t)
	testServer := newTestServer(t, app.RouteHandler())
	defer testServer.Close()

	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{"valid id", "/snip/1", 200, "this is a mock snip"},
		{"zero id", "/snip/0", 404, ErrPageNotFound},
		{"id out of range", "/snip/10", 404, ErrPageNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := testServer.get(t, tt.urlPath)

			// assert status code and body
			testhelpers.AssertInt(t, code, tt.wantCode)
			testhelpers.AssertTemplateContent(t, body, tt.wantBody)
		})
	}
}

func TestApplication_createSnip(t *testing.T) {
	app := newTestApp(t)
	testServer := newTestServer(t, app.RouteHandler())
	defer testServer.Close()

	_, _, body := testServer.get(t, createSnipRoute)
	csrfToken := extractCSRFToken(t, []byte(body))

	tests := []struct {
		name      string
		title     string
		content   string
		csrfToken string
		wantCode  int
		wantBody  string
	}{
		{"valid submission", "test snip", "this is a test snip", csrfToken, http.StatusSeeOther, ""},
		{"missing title", "", "this is a test snip", csrfToken, http.StatusOK, forms.ErrBlankField},
		{"missing content", "test snip", "", csrfToken, http.StatusOK, forms.ErrBlankField},
		{"title above maxlength", "testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttest", "this snip is too long", csrfToken, http.StatusOK, "This field must not have over 100 characters"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{
				forms.Title:     {tt.title},
				forms.Content:   {tt.content},
				forms.CsrfToken: {tt.csrfToken},
			}
			code, _, body := testServer.postForm(t, createSnipRoute, form)

			// assert status code and body
			testhelpers.AssertInt(t, code, tt.wantCode)
			testhelpers.AssertTemplateContent(t, body, tt.wantBody)
		})
	}
}
