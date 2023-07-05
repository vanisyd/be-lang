package api

import (
	"encoding/json"
	"net/http"
	"web/controller/request/language"
	vocabulary2 "web/provider/vocabulary"
	"web/server/kind"
)

func GetLangs() kind.Response {
	var responseContent []byte
	request, valid, errors := language.GetLanguageRequest()

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

	filter := vocabulary2.LanguageFilter()
	for key, value := range request {
		filter[key] = value
	}

	responseContent, _ = vocabulary2.GetLangs(filter)

	return kind.Response{
		StatusCode: http.StatusOK,
		Content:    string(responseContent),
	}
}

func AddLang() kind.Response {
	var responseContent []byte
	/*	request, valid, errors := language.CreateLanguageRequest()

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

		content := vocabulary2.AddLang(vocabulary2.Language{
			ISO: request["iso"].(string),
		})

		responseContent, _ = json.Marshal(content)
	*/
	return kind.Response{
		StatusCode: http.StatusCreated,
		Content:    string(responseContent),
	}
}
