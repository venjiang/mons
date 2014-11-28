package middlewares

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
	. "github.com/venjiang/mons/core"
	"log"
)

// SessionUser will try to read a unique user ID out of the session. Then it tries
// to populate an anonymous user object from the database based on that ID. If this
// is successful, the valid user is mapped into the context. Otherwise the anonymous
// user is mapped into the contact.
// The newUser() function should provide a valid 0value structure for the caller's
// user type.
func Authentication(newAuth func() Authenticator) martini.Handler {
	return func(s sessions.Session, c martini.Context, l *log.Logger) {
		authId := s.Get(SessionKey)
		auth := newAuth()

		if authId != nil {
			err := auth.GetById(authId)
			if err != nil {
				l.Printf("Login Error: %v\n", err)
			} else {
				auth.Login()
			}
		}

		c.MapTo(auth, (*Authenticator)(nil))
	}
}
