package word

import (
	"web/data"
	"web/server"
)

func GetWordRequest() (input map[string]any, err bool) {
	rules := []server.RequestField{
		*server.Rule("id").Sometimes().Int(),
		*server.Rule("language_id").Sometimes().Int(),
		*server.Rule("text"),
		*server.Rule("type").Sometimes().Int(),
		*server.Rule("created_at").Sometimes(),
		*server.Rule(data.KEYWORD_SORT_BY).Sometimes(),
		*server.Rule(data.KEYWORD_SORT_DIR).Sometimes(),
	}

	input, err = server.Validate(rules)

	return
}

func CreateWordRequest() (input map[string]any, err bool) {
	rules := []server.RequestField{
		*server.Rule("language_id").Int(),
		*server.Rule("text").Required(),
		*server.Rule("type").Int(),
		*server.Rule(data.KEYWORD_SORT_BY).Sometimes(),
		*server.Rule(data.KEYWORD_SORT_DIR).Sometimes(),
	}

	input, err = server.Validate(rules)

	return
}
