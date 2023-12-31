package data

const JOIN_INNER = "INNER"
const JOIN_LEFT = "LEFT"
const JOIN_RIGHT = "RIGHT"

const OR = "OR"
const AND = "AND"
const IN = "IN"
const IS_NULL = "IS NULL"
const IS_NOT_NULL = "IS NOT NULL"

const DIR_ASC = "ASC"
const DIR_DESC = "DESC"

const QUERY_TYPE_INSERT = 1
const QUERY_TYPE_SELECT = 2
const QUERY_TYPE_UPDATE = 3
const QUERY_TYPE_DESCRIBE = 4

const KEYWORD_SORT_BY = "sort_by"
const KEYWORD_SORT_DIR = "sort"

const DEFAULT_SORT_DIR = DIR_DESC

func GetSortKeywords() []string {
	return []string{
		KEYWORD_SORT_BY,
		KEYWORD_SORT_DIR,
	}
}
