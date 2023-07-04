package server

import (
	"net/http"
	"net/url"
	"time"
	"web/server/kind"
)

var httpServer = &http.Server{
	Addr:           ":8080",
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
}

var CurrentRoute kind.Route
var HTTPRequest *http.Request
var HTTPQuery url.Values
var HTTPBody map[string]any

func Run(handler http.Handler) {
	httpServer.Handler = handler
	httpServer.ListenAndServe()
}
