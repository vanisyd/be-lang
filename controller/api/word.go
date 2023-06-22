package api

import (
	"fmt"
	"net/http"
	"studying/web/server"
)

func GetWords() server.Response {
	userName := server.HTTPQuery.Get("name")

	return server.Response{
		StatusCode: http.StatusOK,
		Content:    fmt.Sprintf("Hello %s", userName),
	}
}
