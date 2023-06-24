package word

import "studying/web/server"

func GetWordRequest() (input map[string]any, err bool) {
	rules := []server.RequestField{
		*server.Rule("id").Sometimes().Int(),
		*server.Rule("language_id").Sometimes().Int(),
		*server.Rule("text"),
		*server.Rule("type").Sometimes().Int(),
		*server.Rule("created_at").Sometimes(),
		*server.Rule("sort_by").Sometimes(),
	}

	input, err = server.Validate(rules)

	return
}

func CreateWordRequest() (input map[string]any, err bool) {
	rules := []server.RequestField{
		*server.Rule("language_id").Int(),
		*server.Rule("text").Required(),
		*server.Rule("type").Int(),
	}

	input, err = server.Validate(rules)

	return
}
