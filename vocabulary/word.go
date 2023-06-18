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
	ID              int `json:"id"`
	WordID          int `json:"word_id"`
	TranslationID   int `json:"translation_id"`
	Word            Word
	WordTranslation Word
}

var WordModel data.Model = data.Model{
	Name: "words",
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
	Structure: Word{},
}

var LanguageModel data.Model = data.Model{
	Name: "languages",
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
	Name: "translations",
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

func GetWords(lang Language) []Word {
	var words []Word
	dbConnection := data.DBConnection()
	defer dbConnection.Close()

	query := data.Query{
		Model: WordModel,
	}

	query.Select(WordModel, []string{"id", "text", "language_id", "type"})
	query.Select(LanguageModel, []string{"id", "iso"})
	query.Where("words.language_id", "=", fmt.Sprint(lang.ID))
	query.Join(LanguageModel, "language_id", "id")

	rows := data.Get(dbConnection, &query)

	for rows.Next() {
		var dbWord Word
		var dbLang Language
		fmt.Println(rows.Columns())
		rows.Scan(&dbWord.ID, &dbWord.Text, &dbWord.LanguageID, &dbWord.Type, &dbLang.ID, &dbLang.ISO)
		dbWord.Language = dbLang
		words = append(words, dbWord)
	}

	return words
}

func AddWord(wordObj Word) Word {
	dbConnection := data.DBConnection()
	defer dbConnection.Close()

	dbWord, err := dbConnection.Exec("INSERT INTO `words` (`text`, `language_id`, `type`) VALUES (?, ?, ?)", wordObj.Text, wordObj.LanguageID, wordObj.Type)
	if err != nil {
		log.Fatal(err)
		return wordObj
	}
	wordID, err := dbWord.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return wordObj
	}

	wordObj.ID = int(wordID)

	return wordObj
}
