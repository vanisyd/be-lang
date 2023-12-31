package data

import "database/sql"

type ModelField struct {
	Name    string
	SQLName string
}

type Model struct {
	Name    string
	SQLName string
	Columns []ModelField
}

type JoinStmt struct {
	BaseModel    Model
	JoinedModel  Model
	Type         string
	Field        string
	ForeignField string
}

type SelectStmt struct {
	Model Model
}

type UpdateStmt struct {
	Fields map[string]interface{}
}

type ConditionStmt struct {
	Field    string
	Operator string
	Value    string
}

type InsertStmt struct {
	Model  Model
	Values map[string]interface{}
}

type OrderStmt struct {
	Field     string
	Direction string
	Model     Model
}

type Query struct {
	Model          Model
	DBConnection   *sql.DB
	Statements     []interface{}
	SelectStmts    []SelectStmt
	ConditionStmts []ConditionStmt
	JoinStmts      []JoinStmt
	InsertStmt     InsertStmt
	OrderStmts     []OrderStmt
	UpdateStmts    []UpdateStmt
	Data           []map[string]any
	Type           int
}

type Filter map[string]any

func MakeFilter(values map[string]any) Filter {
	newFilter := make(Filter)
	for key, val := range values {
		newFilter[key] = val
	}

	return newFilter
}
