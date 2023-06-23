package api

import (
	"net/http"
	"strconv"
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

	langId, _ := strconv.Atoi(request["lang_id"])
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
