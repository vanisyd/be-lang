package data

import (
	"fmt"
	"log"
)

type TableInfo struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
}

func (model *Model) TableName() string {
	return model.SQLName
}

func (model *Model) GetName() string {
	return fmt.Sprintf("%s", model.Name)
}

func (model *Model) GetColumns() []ModelField {
	return model.Columns
}

func (model *Model) Prepare() {
	if model.GetColumns() == nil {
		var columns []ModelField
		queryString := "SHOW COLUMNS FROM " + model.TableName() + ";"
		rows, err := DBConnection().Query(queryString)
		if err != nil {
			log.Fatal(err)
			return
		}

		for rows.Next() {
			table := TableInfo{}
			rows.Scan(&table.Field, &table.Type, &table.Null, &table.Key, &table.Default, &table.Extra)
			col := ModelField{
				Name:    table.Field,
				SQLName: table.Field,
			}
			columns = append(columns, col)
		}

		model.Columns = columns
	}
}
