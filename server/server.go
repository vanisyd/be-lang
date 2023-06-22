package server

import (
	"net/http"
	"time"
)

var httpServer = &http.Server{
	Addr:           ":8080",
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
}

func Run() {
	httpServer.ListenAndServe()
}
