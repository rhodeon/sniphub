package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/rhodeon/sniphub/cmd/web/internal/templates"
	"github.com/rhodeon/sniphub/pkg/models/mock"
	"github.com/rhodeon/sniphub/pkg/testhelpers"
	"html"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"regexp"
	"testing"
)

// newTestApp generates a test application to provide dependencies for testing.
func newTestApp(t *testing.T) *application {
	t.Helper()

	templateCache, err := templates.NewTemplateCache("./../../ui/html/")
	testhelpers.AssertFatalError(t, err)

	sessionManager := scs.New()
	sessionManager.Cookie.Secure = true
	sessionManager.Cookie.SameSite = http.SameSiteStrictMode

	return &application{
		templateCache:  templateCache,
		sessionManager: sessionManager,
		snips:          &mock.SnipController{},
		users:          &mock.UserController{},
	}
}

// testServer is a wrapper struct for httptest.Server.
type testServer struct {
	*httptest.Server
}

// newTestServer starts and returns a new testServer.
func newTestServer(t *testing.T, h http.Handler) *testServer {
	t.Helper()

	ts := httptest.NewTLSServer(h)

	// set a cookie jar to store cookies
	jar, err := cookiejar.New(nil)
	testhelpers.AssertFatalError(t, err)
	ts.Client().Jar = jar

	// disable redirects
	ts.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return &testServer{ts}
}

// get wraps the Get method of the test server and
// returns the response code, headers and body.
func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, string) {
	t.Helper()

	rs, err := ts.Client().Get(ts.URL + urlPath)
	testhelpers.AssertFatalError(t, err)

	return ts.parseResponse(t, rs)
}

// postForm wraps the PostForm method of the test server and
// returns the response code, headers and body.
func (ts *testServer) postForm(t *testing.T, urlPath string, data url.Values) (int, http.Header, string) {
	t.Helper()

	rs, err := ts.Client().PostForm(ts.URL+urlPath, data)
	testhelpers.AssertFatalError(t, err)

	return ts.parseResponse(t, rs)
}

// parseResponse parses a http response and returns the code, header and body.
func (ts *testServer) parseResponse(t *testing.T, rs *http.Response) (int, http.Header, string) {
	t.Helper()

	body, err := io.ReadAll(rs.Body)
	testhelpers.AssertFatalError(t, err)
	defer rs.Body.Close()
	return rs.StatusCode, rs.Header, string(body)
}

// csrfTokenRX is the pattern of csrf tokens in the HTML templates.
var csrfTokenRX = regexp.MustCompile(`<input name='csrf_token' type='hidden' value='(.+)'>`)

func extractCSRFToken(t *testing.T, body []byte) string {
	t.Helper()

	// Use the FindSubmatch method to extract the token from the HTML body.
	// Note that this returns an array with the entire matched pattern in the
	// first position, and the values of any captured data in the subsequent positions.
	matches := csrfTokenRX.FindSubmatch(body)
	if len(matches) < 2 {
		t.Fatal("no csrf token found in body")
	}
	return html.UnescapeString(string(matches[1]))
}
