package language

import "studying/web/server"

func CreateLanguageRequest() (input map[string]any, err bool) {
	rules := []server.RequestField{
		*server.Rule("iso").Required(),
	}

	input, err = server.Validate(rules)

	return
}

func GetLanguageRequest() (input map[string]any, err bool) {
	rules := []server.RequestField{
		*server.Rule("id").Sometimes().Int(),
		*server.Rule("iso").Sometimes(),
		*server.Rule("created_at").Sometimes(),
	}

	input, err = server.Validate(rules)

	return
}
