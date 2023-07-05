package word

import (
	"web/data"
	"web/provider/vocabulary"
	"web/server/http/validator"
)

func GetWordRequest() (input map[string]any, valid bool, errors map[string][]string) {
	rules := []validator.RequestField{
		*validator.Rule("id").Int(),
		*validator.Rule("language_id").Int().Exists(validator.ExistsRule{
			Model: vocabulary.LanguageModel,
			Field: "id",
		}),
		*validator.Rule("text"),
		*validator.Rule("type").Int(),
		*validator.Rule("created_at"),
		*validator.Rule(data.KEYWORD_SORT_BY),
		*validator.Rule(data.KEYWORD_SORT_DIR).String().In([]string{"asc", "desc"}),
	}

	input, valid, errors = validator.Validate(rules)

	return
}

func CreateWordRequest() (input map[string]any, valid bool, errors map[string][]string) {
	rules := []validator.RequestField{
		*validator.Rule("language_id").Required().Int(),
		*validator.Rule("text").Required().String(),
		*validator.Rule("type").Required().Int(),
	}

	input, valid, errors = validator.Validate(rules)

	return
}

func UpdateWordRequest() (input map[string]any, valid bool, errors map[string][]string) {
	rules := []validator.RequestField{
		*validator.Rule("id").Int(),
		*validator.Rule("language_id").Int().Exists(validator.ExistsRule{
			Model: vocabulary.LanguageModel,
			Field: "id",
		}),
		*validator.Rule("text").String(),
		*validator.Rule("type").Int(),
	}

	input, valid, errors = validator.Validate(rules)

	return
}
