package vocabulary

import "web/data"

func LanguageFilter() data.Filter {
	return data.MakeFilter(map[string]any{
		"id":         nil,
		"iso":        nil,
		"created_at": nil,
	})
}

func WordFilter() data.Filter {
	return data.MakeFilter(map[string]any{
		"id":          nil,
		"text":        nil,
		"language_id": nil,
		"type":        nil,
		"created_at":  nil,
	})
}

var WordModel data.Model = data.Model{
	Name:    "word",
	SQLName: "words",
}

var LanguageModel data.Model = data.Model{
	Name:    "language",
	SQLName: "languages",
}

var TranslationModel data.Model = data.Model{
	Name:    "translation",
	SQLName: "translations",
}

func Init() {
	WordModel.Prepare()
	LanguageModel.Prepare()
	TranslationModel.Prepare()
}
