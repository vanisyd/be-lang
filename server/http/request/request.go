package request

import (
	"strconv"
	"web/server"
)

func GetQueryParam(name string) (param string) {
	param = server.HTTPQuery.Get(name)

	return
}

func GetBodyParam(name string) (param string) {
	if server.HTTPBody != nil {
		neededVal, ok := server.HTTPBody[name]
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
