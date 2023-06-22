package data

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

type ConditionStmt struct {
	Field    string
	Operator string
	Value    string
}

type InsertStmt struct {
	Model  Model
	Values map[string]interface{}
}

type Query struct {
	Model          Model
	Statements     []interface{}
	SelectStmts    []SelectStmt
	ConditionStmts []ConditionStmt
	JoinStmts      []JoinStmt
	InsertStmt     InsertStmt
	Data           []map[string]any
	Type           int
}
