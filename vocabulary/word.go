package vocabulary

import (
	"log"
	"web/data"
)

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
		Model:        WordModel,
		DBConnection: dbConnection,
	}

	query.Select(WordModel).Select(LanguageModel)
	query.Join(WordModel, LanguageModel, "language_id", "id")
	query.Filter(filter)
	query.Get()

	return query.ToJson()
}

func GetLangs(filter data.Filter) ([]byte, error) {
	dbConnection := data.DBConnection()
	query := data.Query{
		Model:        LanguageModel,
		DBConnection: dbConnection,
	}

	query.Select(LanguageModel)
	query.Filter(filter)
	query.Get()

	return query.ToJson()
}

func AddWord(input map[string]interface{}) int64 {
	dbConnection := data.DBConnection()

	query := data.Query{
		Model:        WordModel,
		DBConnection: dbConnection,
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

func UpdateWord(input map[string]interface{}, filter data.Filter) int64 {
	dbConnection := data.DBConnection()

	query := data.Query{
		Model:        WordModel,
		DBConnection: dbConnection,
	}

	query.Update(input)
	query.Filter(filter)
	result := query.Exec(dbConnection)

	rowsNumber, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return 0
	}

	return rowsNumber
}

func AddLang(langObj Language) Language {
	dbConnection := data.DBConnection()

	query := data.Query{
		Model:        LanguageModel,
		DBConnection: dbConnection,
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
