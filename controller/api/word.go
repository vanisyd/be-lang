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

	filter := vocabulary.WordFilter
	for key, value := range request {
		filter[key] = value
	}

	content, _ := vocabulary.GetWords(filter)

	return server.Response{
		StatusCode: http.StatusOK,
		Content:    string(content),
	}
}

func AddWord() server.Response {
	request, valid := word.CreateWordRequest()
	if !valid {
		return server.Response{
			StatusCode: http.StatusUnprocessableEntity,
		}
	}

	var wordData []byte
	wordId := vocabulary.AddWord(request)
	if wordId != 0 {
		filter := vocabulary.WordFilter
		filter["id"] = wordId
		wordData, _ = vocabulary.GetWords(filter)
	}

	return server.Response{
		StatusCode: http.StatusAccepted,
		Content:    string(wordData),
	}
}
