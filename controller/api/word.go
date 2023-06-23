package api

import (
	"net/http"
	"studying/web/controller/request/word"
	"studying/web/server"
	"studying/web/vocabulary"
)

func GetWords() server.Response {
	request, valid := word.GetWordRequest()

	if !valid {
		return server.Response{
			StatusCode: http.StatusUnprocessableEntity,
		}
	}

	content, _ := vocabulary.GetWords(request["lang_id"].(int))

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
