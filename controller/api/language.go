package api

import (
	"encoding/json"
	"net/http"
	"web/controller/request/language"
	"web/server"
	"web/vocabulary"
)

func GetLangs() server.Response {
	var responseContent []byte
	request, valid, errors := language.GetLanguageRequest()

	if !valid {
		errorsJson, err := json.Marshal(errors)
		if err == nil {
			responseContent = errorsJson
		}
		return server.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Content:    string(responseContent),
		}
	}

	filter := vocabulary.LanguageFilter()
	for key, value := range request {
		filter[key] = value
	}

	responseContent, _ = vocabulary.GetLangs(filter)

	return server.Response{
		StatusCode: http.StatusOK,
		Content:    string(responseContent),
	}
}

func AddLang() server.Response {
	var responseContent []byte
	request, valid, errors := language.CreateLanguageRequest()

	if !valid {
		errorsJson, err := json.Marshal(errors)
		if err == nil {
			responseContent = errorsJson
		}
		return server.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Content:    string(responseContent),
		}
	}

	content := vocabulary.AddLang(vocabulary.Language{
		ISO: request["iso"].(string),
	})

	responseContent, _ = json.Marshal(content)

	return server.Response{
		StatusCode: http.StatusCreated,
		Content:    string(responseContent),
	}
}
