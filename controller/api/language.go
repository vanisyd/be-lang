package api

import (
	"encoding/json"
	"net/http"
	"studying/web/controller/request/language"
	"studying/web/server"
	"studying/web/vocabulary"
)

func AddLang() server.Response {
	request, valid := language.CreateLanguageRequest()

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
