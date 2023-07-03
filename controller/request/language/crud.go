package language

import "web/server"

func CreateLanguageRequest() (input map[string]any, valid bool, errors map[string][]string) {
	rules := []server.RequestField{
		*server.Rule("iso").Required(),
	}

	input, valid, errors = server.Validate(rules)

	return
}

func GetLanguageRequest() (input map[string]any, valid bool, errors map[string][]string) {
	rules := []server.RequestField{
		*server.Rule("id").Sometimes().Int(),
		*server.Rule("iso").Sometimes(),
		*server.Rule("created_at").Sometimes(),
	}

	input, valid, errors = server.Validate(rules)

	return
}
