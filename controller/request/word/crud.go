package word

import (
	"web/data"
	"web/server"
)

func GetWordRequest() (input map[string]any, valid bool, errors map[string][]string) {
	rules := []server.RequestField{
		*server.Rule("id").Sometimes().Int(),
		*server.Rule("language_id").Sometimes().Int(),
		*server.Rule("text").Sometimes(),
		*server.Rule("type").Sometimes().Int(),
		*server.Rule("created_at").Sometimes(),
		*server.Rule(data.KEYWORD_SORT_BY).Sometimes(),
		*server.Rule(data.KEYWORD_SORT_DIR).Sometimes(),
	}

	input, valid, errors = server.Validate(rules)

	return
}

func CreateWordRequest() (input map[string]any, valid bool, errors map[string][]string) {
	rules := []server.RequestField{
		*server.Rule("language_id").Int(),
		*server.Rule("text").Required(),
		*server.Rule("type").Int(),
	}

	input, valid, errors = server.Validate(rules)

	return
}

func UpdateWordRequest() (input map[string]any, valid bool, errors map[string][]string) {
	rules := []server.RequestField{
		*server.Rule("id").Int(),
		*server.Rule("language_id").Sometimes().Int(),
		*server.Rule("text").Sometimes(),
		*server.Rule("type").Sometimes().Int(),
	}

	input, valid, errors = server.Validate(rules)

	return
}
