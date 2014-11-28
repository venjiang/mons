package core

const (
	UserNotFound = "UserNotFound"
	ServerError  = "ServerError"
)

type AppErr struct {
	Type    string
	Message string
}
