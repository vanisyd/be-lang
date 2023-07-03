package word

import (
	"web/data"
	"web/server"
)

func GetWordRequest() (input map[string]any, valid bool, errors map[string][]string) {
	rules := []server.RequestField{
		*server.Rule("id").Int(),
		*server.Rule("language_id").Int(),
		*server.Rule("text"),
		*server.Rule("type").Int(),
		*server.Rule("created_at"),
		*server.Rule(data.KEYWORD_SORT_BY),
		*server.Rule(data.KEYWORD_SORT_DIR).String().In([]string{"asc", "desc"}),
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
		*server.Rule("language_id").Int(),
		*server.Rule("text"),
		*server.Rule("type").Int(),
	}

	input, valid, errors = server.Validate(rules)

	return
}
