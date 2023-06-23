package language

import "studying/web/server"

func CreateLanguageRequest() (input map[string]any, err bool) {
	rules := []server.RequestField{
		*server.Rule("iso").Required(),
	}

	input, err = server.Validate(rules)

	return
}
