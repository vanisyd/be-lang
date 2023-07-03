package api

import (
	"encoding/json"
	"net/http"
	"web/controller/request/language"
	"web/server"
	"web/vocabulary"
)

func GetLangs() server.Response {
	request, valid, _ := language.GetLanguageRequest()

	if !valid {
		return server.Response{
			StatusCode: http.StatusUnprocessableEntity,
		}
	}

	filter := vocabulary.LanguageFilter()
	for key, value := range request {
		filter[key] = value
	}

	content, _ := vocabulary.GetLangs(filter)

	return server.Response{
		StatusCode: http.StatusOK,
		Content:    string(content),
	}
}

func AddLang() server.Response {
	request, valid, _ := language.CreateLanguageRequest()

	if !valid {
		return server.Response{
			StatusCode: http.StatusUnprocessableEntity,
		}
	}

	content := vocabulary.AddLang(vocabulary.Language{
		ISO: request["iso"].(string),
	})

	jsonContent, _ := json.Marshal(content)

	return server.Response{
		StatusCode: http.StatusCreated,
		Content:    string(jsonContent),
	}
}
