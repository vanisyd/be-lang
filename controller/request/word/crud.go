package word

import "studying/web/server"

func GetWordRequest() (input map[string]string, err bool) {
	rules := []server.RequestField{
		*server.Rule("lang_id").Required(),
		*server.Rule("user"),
	}

	input, err = server.Validate(rules)

	return
}
