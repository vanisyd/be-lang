package data

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"studying/web/config"
)

var DB *sql.DB

type ModelField struct {
	Name    string
	SQLName string
}

type Model struct {
	Name      string
	Columns   []ModelField
	Structure interface{}
}

type JoinStmt struct {
	Model        Model
	Type         string
	Field        string
	ForeignField string
}

type SelectStmt struct {
	Model  Model
	Fields []string
}

type ConditionStmt struct {
	Field    string
	Operator string
	Value    string
}

type Query struct {
	Model          Model
	Statements     []interface{}
	SelectStmts    []SelectStmt
	ConditionStmts []ConditionStmt
	JoinStmts      []JoinStmt
}

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

func (query *Query) Select(model Model, fields []string) *Query {
	query.AddStmt(SelectStmt{
		model,
		fields,
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

func (query *Query) Join(model Model, field string, foreignField string) *Query {
	query.AddStmt(JoinStmt{
		model,
		JOIN_INNER,
		field,
		foreignField,
	})

	return query
}

func (model *Model) TableName() string {
	return model.Name
}

func DBConnection() *sql.DB {
	if DB == nil {
		var err error
		DB, err = sql.Open("mysql", config.DBConfig.FormatDSN())

		if err != nil {
			log.Fatal(err)
		}

		pingErr := DB.Ping()
		if pingErr != nil {
			log.Fatal(pingErr)
		}

		fmt.Println("[DB] Connected")
	}

	return DB
}

func Get(dbConnection *sql.DB, queryObj *Query) (rows *sql.Rows) {
	query := prepareQuery(queryObj)
	rows, err := dbConnection.Query(query)
	if err != nil {
		log.Fatal(err)
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
		columns := prepareQueryColumns(stmt.Model, stmt.Fields)
		queryString += columns
	}
	queryString += fmt.Sprintf(" FROM %s", query.Model.TableName())
	for _, stmt := range query.JoinStmts {
		queryString += " " + fmt.Sprintf("%s JOIN %s ON %s.%s = %s.%s", stmt.Type, stmt.Model.TableName(), query.Model.TableName(), stmt.Field, stmt.Model.TableName(), stmt.ForeignField)
	}
	for i, stmt := range query.ConditionStmts {
		if i == 0 {
			queryString += " WHERE"
		}

		queryString += " " + fmt.Sprintf("%s %s %s", stmt.Field, stmt.Operator, stmt.Value)
	}

	return
}
