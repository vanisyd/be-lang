package vocabulary

import (
	"fmt"
	"log"
	"web/data"
	"web/data/model"
	"web/dataprovider/querybuilder"
	"web/provider/authservice"
)

func GetWords(filter data.Filter) ([]byte, error) {
	var userRes model.Resource
	query := querybuilder.Query{
		Model: WordModel,
	}
	result := query.Select(authservice.UserModel).Parse(userRes)
	fmt.Printf("%v", result)
	/*	query := data.Query{
			Model:        WordModel,
			DBConnection: dbConnection,
		}

		query.Select(WordModel).Select(LanguageModel)
		query.Join(WordModel, LanguageModel, "language_id", "id")
		query.Filter(filter)
		query.Get()*/

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
