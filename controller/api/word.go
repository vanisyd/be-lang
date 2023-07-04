package api

import (
	"encoding/json"
	"net/http"
	"web/controller/request/word"
	"web/helper"
	"web/server/kind"
	"web/vocabulary"
)

func GetWords() kind.Response {
	var responseContent []byte
	request, valid, errors := word.GetWordRequest()

	if !valid {
		errorsJson, err := json.Marshal(errors)
		if err == nil {
			responseContent = errorsJson
		}
		return kind.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Content:    string(responseContent),
		}

	}

	filter := vocabulary.WordFilter()
	for key, value := range request {
		filter[key] = value
	}

	responseContent, _ = vocabulary.GetWords(filter)

	return kind.Response{
		StatusCode: http.StatusOK,
		Content:    string(responseContent),
	}
}

func AddWord() kind.Response {
	var responseContent []byte
	request, valid, errors := word.CreateWordRequest()

	if !valid {
		errorsJson, err := json.Marshal(errors)
		if err == nil {
			responseContent = errorsJson
		}
		return kind.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Content:    string(responseContent),
		}
	}

	wordId := vocabulary.AddWord(request)
	if wordId != 0 {
		filter := vocabulary.WordFilter()
		filter["id"] = wordId
		responseContent, _ = vocabulary.GetWords(filter)
	}

	return kind.Response{
		StatusCode: http.StatusCreated,
		Content:    string(responseContent),
	}
}

func UpdateWord() kind.Response {
	var responseContent []byte
	request, valid, errors := word.UpdateWordRequest()

	if !valid {
		errorsJson, err := json.Marshal(errors)
		if err == nil {
			responseContent = errorsJson
		}
		return kind.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Content:    string(responseContent),
		}
	}

	filter := vocabulary.WordFilter()
	filter["id"] = request["id"]
	result := vocabulary.UpdateWord(helper.Except(request, []string{"id"}), filter)
	if result != 0 {
		responseContent, _ = vocabulary.GetWords(filter)
	}

	return kind.Response{
		StatusCode: http.StatusOK,
		Content:    string(responseContent),
	}
}
