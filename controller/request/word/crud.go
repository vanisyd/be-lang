package word

import "studying/web/server"

func GetWordRequest() (input map[string]any, err bool) {
	rules := []server.RequestField{
		*server.Rule("lang_id").Required().Int(),
		*server.Rule("user"),
	}

	input, err = server.Validate(rules)

	return
}
