package middleware

import (
	"fmt"
	"net/http"
	"web/server"
	"web/server/kind"
)

const MIDDLEWARE_STATUS_BREAK int = 0
const MIDDLEWARE_STATUS_CONTINUE int = 1

type AuthMiddleware struct {
	Status int
}

var Auth kind.HTTPMiddleware = AuthMiddleware{}

func (middleware AuthMiddleware) Handle(w http.ResponseWriter) kind.HTTPMiddleware {
	authHeader, ok := server.HTTPRequest.Header["Authorization"]
	if ok {
		fmt.Println(authHeader)
		middleware.Status = MIDDLEWARE_STATUS_CONTINUE
	} else {
		w.WriteHeader(http.StatusForbidden)
		middleware.Status = MIDDLEWARE_STATUS_BREAK
	}

	return middleware
}

func (middleware AuthMiddleware) GetStatus() int {
	return middleware.Status
}
