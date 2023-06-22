package router

type Route struct {
	Path    string
	Handler func()
}
