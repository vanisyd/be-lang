package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func (query *Query) AddStmt(stmt interface{}) *Query {
	switch statement := stmt.(type) {
	case SelectStmt:
		query.SelectStmts = append(query.SelectStmts, statement)
	case ConditionStmt:
		query.ConditionStmts = append(query.ConditionStmts, statement)
	case JoinStmt:
		query.JoinStmts = append(query.JoinStmts, statement)
	}

	return query
}

func (query *Query) Select(model Model) *Query {
	query.AddStmt(SelectStmt{
		model,
	})

	return query
}

func (query *Query) Where(field string, operator string, value string) *Query {
	query.AddStmt(ConditionStmt{
		field,
		operator,
		value,
	})

	return query
}

func (query *Query) Join(baseModel Model, model Model, field string, foreignField string) *Query {
	query.AddStmt(JoinStmt{
		baseModel,
		model,
		JOIN_INNER,
		field,
		foreignField,
	})

	return query
}

func (query *Query) ToJson() ([]byte, error) {
	return json.Marshal(query.Data)
}

func (query *Query) clearData() {
	query.Data = make([]map[string]any, 0)
}

func (stmt *SelectStmt) getColumns() (columns []string) {
	for _, column := range stmt.Model.Columns {
		columns = append(columns, column.SQLName)
	}

	return columns
}

func Get(dbConnection *sql.DB, queryObj *Query) {
	query := prepareQuery(queryObj)
	rows, err := dbConnection.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	columns, _ := rows.Columns()
	data := make([][]byte, len(columns))
	dataPtr := make([]any, len(columns))

	for i := range data {
		dataPtr[i] = &data[i]
	}

	for rows.Next() {
		rows.Scan(dataPtr...)

		serializedData := make(map[string]interface{}, len(columns))
		colIndex := 0
		for _, stmt := range queryObj.SelectStmts {
			serializedStmt := make(map[string]interface{}, len(stmt.Model.Columns))
			for _, column := range stmt.Model.Columns {
				value := string(data[colIndex])
				serializedStmt[column.Name] = value
				colIndex++
			}

			if stmt.Model.GetName() == queryObj.Model.GetName() {
				for key, val := range serializedStmt {
					serializedData[key] = val
				}
			} else {
				serializedData[stmt.Model.GetName()] = map[string]interface{}{}
				for key, val := range serializedStmt {
					nestedData, ok := serializedData[stmt.Model.GetName()].(map[string]any)
					if ok == true {
						nestedData[key] = val
					}
				}
			}
		}
		queryObj.Data = append(queryObj.Data, serializedData)
	}

	return
}

func prepareQuery(queryObj *Query) (query string) {
	query = prepareStatements(queryObj)

	return
}

func prepareQueryColumns(model Model, fields []string) (columns string) {
	var columnsArray []string
	for _, column := range fields {
		columnsArray = append(columnsArray, model.TableName()+"."+column)
	}
	columns = strings.Join(columnsArray, ",")

	return
}

func prepareStatements(query *Query) (queryString string) {
	if len(query.SelectStmts) > 0 {
		queryString += "SELECT"
	}

	for i, stmt := range query.SelectStmts {
		if i != 0 {
			queryString += ","
		} else {
			queryString += " "
		}
		columns := prepareQueryColumns(stmt.Model, stmt.getColumns())
		queryString += columns
	}
	queryString += fmt.Sprintf(" FROM %s", query.Model.TableName())
	for _, stmt := range query.JoinStmts {
		queryString += " " + fmt.Sprintf("%s JOIN %s ON %s.%s = %s.%s", stmt.Type, stmt.JoinedModel.TableName(), stmt.BaseModel.TableName(), stmt.Field, stmt.JoinedModel.TableName(), stmt.ForeignField)
	}
	for i, stmt := range query.ConditionStmts {
		if i == 0 {
			queryString += " WHERE"
		}

		queryString += " " + fmt.Sprintf("%s %s %s", stmt.Field, stmt.Operator, stmt.Value)
	}

	return
}
