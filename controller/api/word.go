package api

import (
	"net/http"
	"web/controller/request/word"
	"web/helper"
	"web/server"
	"web/vocabulary"
)

func GetWords() server.Response {
	request, valid, _ := word.GetWordRequest()

	if !valid {
		return server.Response{
			StatusCode: http.StatusUnprocessableEntity,
		}
	}

	filter := vocabulary.WordFilter()
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
	request, valid, _ := word.CreateWordRequest()
	if !valid {
		return server.Response{
			StatusCode: http.StatusUnprocessableEntity,
		}
	}

	var wordData []byte
	wordId := vocabulary.AddWord(request)
	if wordId != 0 {
		filter := vocabulary.WordFilter()
		filter["id"] = wordId
		wordData, _ = vocabulary.GetWords(filter)
	}

	return server.Response{
		StatusCode: http.StatusCreated,
		Content:    string(wordData),
	}
}

func UpdateWord() server.Response {
	request, valid, _ := word.UpdateWordRequest()
	if !valid {
		return server.Response{
			StatusCode: http.StatusUnprocessableEntity,
		}
	}

	var wordData []byte
	filter := vocabulary.WordFilter()
	filter["id"] = request["id"]
	result := vocabulary.UpdateWord(helper.Except(request, []string{"id"}), filter)
	if result != 0 {
		wordData, _ = vocabulary.GetWords(filter)
	}

	return server.Response{
		StatusCode: http.StatusOK,
		Content:    string(wordData),
	}
}
