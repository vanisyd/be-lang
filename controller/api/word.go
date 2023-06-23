package api

import (
	"net/http"
	"strconv"
	"studying/web/server"
	"studying/web/vocabulary"
)

func GetWords() server.Response {
	langId, err := strconv.Atoi(server.HTTPQuery.Get("lang_id"))
	if err != nil {
		return server.Response{
			StatusCode: http.StatusUnprocessableEntity,
		}
	}

	content, _ := vocabulary.GetWords(langId)

	return server.Response{
		StatusCode: http.StatusOK,
		Content:    string(content),
	}
}

func AddWord() server.Response {
	return server.Response{
		StatusCode: http.StatusAccepted,
		Content:    "works",
	}
}
