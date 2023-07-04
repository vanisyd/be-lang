package language

import "web/server/http/validator"

func CreateLanguageRequest() (input map[string]any, valid bool, errors map[string][]string) {
	rules := []validator.RequestField{
		*validator.Rule("iso").Required(),
	}

	input, valid, errors = validator.Validate(rules)

	return
}

func GetLanguageRequest() (input map[string]any, valid bool, errors map[string][]string) {
	rules := []validator.RequestField{
		*validator.Rule("id").Int(),
		*validator.Rule("iso"),
		*validator.Rule("created_at"),
	}

	input, valid, errors = validator.Validate(rules)

	return
}
