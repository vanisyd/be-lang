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
	w.WriteHeader(resp.StatusCode)
	fmt.Fprintf(w, "%s", resp.Content)
}
