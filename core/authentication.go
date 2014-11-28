package core

import (
	"fmt"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
)

// These are the default configuration values for this package. They
// can be set at anytime, probably during the initial setup of Martini.
var (
	// RedirectUrl should be the relative URL for your login route
	LoginUrl string = "/login"

	// DefaultUrl is login success rerurn url
	DefaultUrl string = "/"

	// RedirectParam is the query string parameter that will be set
	// with the page the auth was trying to visit before they were
	// intercepted.
	RedirectParam string = "next"

	// Key is the key containing the unique ID in your authentication
	SessionKey string = "AUTHUNIQUEID"
)

// Authenticate will mark the session and auth object as authenticated. Then
// the Login() auth function will be called. This function should be called after
// you have validated a auth.
func Authenticate(s sessions.Session, auth Authenticator, remebmer ...bool) error {
	auth.Login()
	if len(remebmer) > 0 && remebmer[0] {
		s.Options(sessions.Options{MaxAge: 60 * 60 * 24 * 30})
	}
	return UpdateUser(s, auth)
}
func AuthenticateRedirect(s sessions.Session, auth Authenticator, r render.Render, req *http.Request, remebmer ...bool) {
	auth.Login()
	if len(remebmer) > 0 && remebmer[0] {
		s.Options(sessions.Options{MaxAge: 60 * 60 * 24 * 30})
	}
	UpdateUser(s, auth)

	if auth.IsAuthenticated() {
		path := req.URL.Query().Get(RedirectParam)
		if len(path) > 0 {
			r.Redirect(path, 302)
			return
		}
		r.Redirect(DefaultUrl, 302)
	}
}

// Logout will clear out the session and call the Logout() auth function.
func Logout(s sessions.Session, auth Authenticator) {
	auth.Logout()
	s.Delete(SessionKey)
}

// LoginRequired verifies that the current auth is authenticated. Any routes that
// require a login should have this handler placed in the flow. If the auth is not
// authenticated, they will be redirected to /login with the "next" get parameter
// set to the attempted URL.
func LoginRequired(r render.Render, auth Authenticator, req *http.Request) {
	if auth.IsAuthenticated() == false {
		path := fmt.Sprintf("%s?%s=%s", LoginUrl, RedirectParam, req.URL.Path)
		r.Redirect(path, 302)
	}
}

// UpdateUser updates the Authenticator object stored in the session. This is useful incase a change
// is made to the auth model that needs to persist across requests.
func UpdateUser(s sessions.Session, auth Authenticator) error {
	s.Set(SessionKey, auth.UniqueId())
	return nil
}
