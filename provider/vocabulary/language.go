package vocabulary

import (
	"web/data"
)

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

/*func AddLang(langObj Language) Language {
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
}*/
