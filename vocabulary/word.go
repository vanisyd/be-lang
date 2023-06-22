package vocabulary

import (
	"fmt"
	"log"
	"studying/web/data"
)

type Word struct {
	ID         int    `json:"id"`
	Text       string `json:"text"`
	LanguageID int    `json:"language_id"`
	Type       int    `json:"type"`
	CreatedAt  string `json:"created_at"`
	Language   Language
}

type Language struct {
	ID        int    `json:"id"`
	ISO       string `json:"iso"`
	CreatedAt string `json:"created_at"`
}

type Translation struct {
	ID            int `json:"id"`
	WordID        int `json:"word_id"`
	TranslationID int `json:"translation_id"`
}

var WordModel data.Model = data.Model{
	Name:    "Word",
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
	Name:    "Language",
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
	Name:    "Translation",
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

func (*Word) MakeArray() []Word {
	return []Word{}
}

func (*Word) MakeInstance() Word {
	return Word{}
}

func GetWords(lang Language) ([]byte, error) {
	dbConnection := data.DBConnection()
	defer dbConnection.Close()

	query := data.Query{
		Model: WordModel,
	}

	query.Select(WordModel).Select(LanguageModel)
	query.Join(WordModel, LanguageModel, "language_id", "id")
	query.Where("words.language_id", "=", fmt.Sprint(lang.ID))
	data.Get(dbConnection, &query)

	return query.ToJson()
}

func AddWord(wordObj Word) Word {
	dbConnection := data.DBConnection()
	defer dbConnection.Close()

	query := data.Query{
		Model: WordModel,
	}

	query.Insert(map[string]interface{}{
		"text":        wordObj.Text,
		"language_id": wordObj.LanguageID,
		"type":        wordObj.Type,
	})
	result := data.Execute(dbConnection, &query)

	wordID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return wordObj
	}

	wordObj.ID = int(wordID)

	return wordObj
}
