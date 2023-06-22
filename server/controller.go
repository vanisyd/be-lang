package server

import (
	"fmt"
	"net/http"
	"net/url"
)

type Action func() Response

type Response struct {
	StatusCode int
	Content    string
}

type Request struct {
	Values map[string]any
}

var HTTPRequest *http.Request
var HTTPQuery url.Values

func GetResponse(w http.ResponseWriter, r *http.Request, handler Action) {
	HTTPRequest = r
	HTTPQuery = HTTPRequest.URL.Query()
	resp := handler()
	fmt.Fprintf(w, "[%d] %s", resp.StatusCode, resp.Content)
}
