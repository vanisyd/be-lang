package server

type Route struct {
	Path    string
	Handler Action
	Method  string
}
