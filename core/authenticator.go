package core

// User defines all the functions necessary to work with the user's authentication.
// The caller should implement these functions for whatever system of authentication
// they choose to use
type Authenticator interface {
	// Return whether this user is logged in or not
	IsAuthenticated() bool

	// Set any flags or extra data that should be available
	Login()

	// Clear any sensitive data out of the user
	Logout()

	// Return the unique identifier of this user object
	UniqueId() interface{}

	// Populate this user object with values
	GetById(id interface{}) error
}
