package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Action func() Response

type Response struct {
	StatusCode int
	Content    string
}

type Request struct {
	Values map[string]any
}

var HTTPRequest *http.Request
var HTTPQuery url.Values
var HTTPBody map[string]any

func GetResponse(w http.ResponseWriter, r *http.Request, handler Action) {
	HTTPRequest = r
	HTTPQuery = HTTPRequest.URL.Query()
	json.NewDecoder(HTTPRequest.Body).Decode(&HTTPBody)
	resp := handler()
	w.WriteHeader(resp.StatusCode)
	fmt.Fprintf(w, "%s", resp.Content)
}

func GetQueryParam(name string) (param string) {
	param = HTTPQuery.Get(name)

	return
}

func GetBodyParam(name string) (param string) {
	if HTTPBody != nil {
		neededVal, ok := HTTPBody[name]
		if ok {
			switch convVal := neededVal.(type) {
			case int:
				param = strconv.Itoa(convVal)
			case float64:
				param = strconv.FormatFloat(convVal, 'G', -1, 32)
			case string:
				param = convVal
			}
		}
	}

	return
}

func GetParam(name string) (param string) {
	param = GetQueryParam(name)
	if param == "" {
		param = GetBodyParam(name)
	}

	return
}
