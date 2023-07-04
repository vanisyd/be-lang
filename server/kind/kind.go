package kind

import "net/http"

type Request struct {
	Values map[string]any
}

type Response struct {
	StatusCode int
	Content    string
}

type Route struct {
	Path        string
	Handler     func() Response
	Method      string
	Middlewares []HTTPMiddleware
	Children    []Route
}

type RouterMapping map[string]interface{}
type RouterMappingItem struct {
	Route    Route
	Children RouterMapping
}

type HTTPMiddleware interface {
	Handle(w http.ResponseWriter) HTTPMiddleware
	GetStatus() int
}
