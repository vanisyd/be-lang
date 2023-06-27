package vocabulary

import (
	"log"
	"web/data"
)

var WordFilter data.Filter = data.Filter{
	"id":          nil,
	"text":        nil,
	"language_id": nil,
	"type":        nil,
	"created_at":  nil,
}

var LanguageFilter data.Filter = data.Filter{
	"id":         nil,
	"iso":        nil,
	"created_at": nil,
}

var WordModel data.Model = data.Model{
	Name:    "word",
	SQLName: "words",
	Columns: []data.ModelField{
		{
			Name:    "ID",
			SQLName: "id",
		},
		{
			Name:    "Text",
			SQLName: "text",
		},
		{
			Name:    "LanguageID",
			SQLName: "language_id",
		},
		{
			Name:    "Type",
			SQLName: "type",
		},
		{
			Name:    "CreatedAt",
			SQLName: "created_at",
		},
	},
}

var LanguageModel data.Model = data.Model{
	Name:    "language",
	SQLName: "languages",
	Columns: []data.ModelField{
		{
			Name:    "ID",
			SQLName: "id",
		},
		{
			Name:    "ISO",
			SQLName: "iso",
		},
		{
			Name:    "CreatedAt",
			SQLName: "created_at",
		},
	},
}

var TranslationModel data.Model = data.Model{
	Name:    "translation",
	SQLName: "translations",
	Columns: []data.ModelField{
		{
			Name:    "ID",
			SQLName: "id",
		},
		{
			Name:    "WordId",
			SQLName: "word_id",
		},
		{
			Name:    "TranslationId",
			SQLName: "translation_id",
		},
	},
}

func GetWords(filter data.Filter) ([]byte, error) {
	dbConnection := data.DBConnection()
	query := data.Query{
		Model: WordModel,
	}

	query.Select(WordModel).Select(LanguageModel)
	query.Join(WordModel, LanguageModel, "language_id", "id")
	query.Filter(filter)
	query.Get(dbConnection)

	return query.ToJson()
}

func GetLangs(filter data.Filter) ([]byte, error) {
	dbConnection := data.DBConnection()
	query := data.Query{
		Model: LanguageModel,
	}

	query.Select(LanguageModel)
	query.Filter(filter)
	query.Get(dbConnection)

	return query.ToJson()
}

func AddWord(input map[string]interface{}) int64 {
	dbConnection := data.DBConnection()

	query := data.Query{
		Model: WordModel,
	}

	query.Insert(input)
	result := query.Exec(dbConnection)

	wordID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return 0
	}

	return wordID
}

func AddLang(langObj Language) Language {
	dbConnection := data.DBConnection()

	query := data.Query{
		Model: LanguageModel,
	}

	query.Insert(map[string]interface{}{
		"iso": langObj.ISO,
	})
	result := query.Exec(dbConnection)

	langID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	langObj.ID = int(langID)

	return langObj
}
