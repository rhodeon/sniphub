package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/rhodeon/sniphub/pkg/models/mock"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"
)

// newTestApp generates a test application to provide dependencies for testing.
func newTestApp(t *testing.T) *application {
	t.Helper()

	templateCache, err := newTemplateCache("./../../ui/html/")
	if err != nil {
		t.Fatal(err)
	}

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

	ts := httptest.NewServer(h)

	// set a cookie jar to store cookies
	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}
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
	rs, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Body.Close()

	return rs.StatusCode, rs.Header, string(body)
}
