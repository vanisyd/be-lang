package validator

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"web/data"
	"web/helper"
)

// Base rule
type ValidationRule interface {
	Check(field *RequestField, errors *map[string][]string) (valid bool)
}

type RequestField struct {
	Name           string
	Rules          []ValidationRule
	Value          string
	ConvertedValue interface{}
}

// Required
type RequiredRule struct{}

func (rule RequiredRule) Check(field *RequestField, errors *map[string][]string) (valid bool) {
	valid = true

	if field.Value == "" {
		addError(field.Name, errors, VALIDATION_ERROR_REQUIRED, &valid)
	}

	return
}

func (field *RequestField) Required() *RequestField {
	var rule ValidationRule = RequiredRule{}
	field.Rules = append(field.Rules, rule)

	return field
}

// Type
type TypeRule struct {
	FieldType string
}

func (rule TypeRule) Check(field *RequestField, errors *map[string][]string) (valid bool) {
	valid = true
	if field.Value == "" {
		return
	}

	switch rule.FieldType {
	case TYPE_INT:
		intValue, error := strconv.Atoi(field.Value)
		if error == nil {
			field.ConvertedValue = intValue
		} else {
			addError(field.Name, errors, fmt.Sprintf(VALIDATION_ERROR_TYPE, TYPE_INT), &valid)
		}
	case TYPE_STRING:
		field.ConvertedValue = field.Value
	}

	return
}

func (field *RequestField) Int() *RequestField {
	var rule ValidationRule = TypeRule{
		FieldType: TYPE_INT,
	}
	field.Rules = append(field.Rules, rule)

	return field
}

func (field *RequestField) String() *RequestField {
	var rule ValidationRule = TypeRule{
		FieldType: TYPE_STRING,
	}
	field.Rules = append(field.Rules, rule)

	return field
}

// In
type InRule[T string | int] struct {
	Values []T
}

func (rule InRule[T]) Check(field *RequestField, errors *map[string][]string) (valid bool) {
	valid = true
	if field.Value == "" {
		return
	}

	if !helper.InSlice(rule.Values, field.ConvertedValue.(T)) {
		addError(field.Name, errors, VALIDATION_ERROR_IN, &valid)
	}

	return
}

func (field *RequestField) In(values interface{}) *RequestField {
	if reflect.TypeOf(values).Kind() != reflect.Slice {
		log.Fatal("Incorrect parameter type, attribute must be a slice")
	}

	var rule ValidationRule

	switch convValues := values.(type) {
	case []string:
		rule = InRule[string]{
			Values: convValues,
		}
	case []int:
		rule = InRule[int]{
			Values: convValues,
		}
	}

	field.Rules = append(field.Rules, rule)

	return field
}

// Exists
type ExistsRule struct {
	Model data.Model
	Field string
}

func (rule ExistsRule) Check(field *RequestField, errors *map[string][]string) (valid bool) {
	valid = true
	if field.Value == "" {
		return
	}

	dbConnection := data.DBConnection()
	query := data.Query{
		Model:        rule.Model,
		DBConnection: dbConnection,
	}

	query.Select(rule.Model).Where(rule.Field, "=", field.Value).Get()

	if len(query.Data) == 0 {
		addError(field.Name, errors, fmt.Sprintf(VALIDATION_ERROR_EXISTS, field.Value), &valid)
	}

	return
}

func (field *RequestField) Exists(newRule ExistsRule) *RequestField {
	var rule ValidationRule = newRule

	field.Rules = append(field.Rules, rule)

	return field
}

// Unique
type UniqueRule struct {
	Model       data.Model
	Field       string
	IgnoreValue string
}

func (rule UniqueRule) Check(field *RequestField, errors *map[string][]string) (valid bool) {
	valid = true
	if field.Value == "" {
		return
	}

	dbConnection := data.DBConnection()
	query := data.Query{
		Model:        rule.Model,
		DBConnection: dbConnection,
	}

	query.Select(rule.Model).Where(rule.Field, "=", field.Value)
	if rule.IgnoreValue != "" {
		query.Where(rule.Field, "!=", rule.IgnoreValue)
	}
	query.Get()

	if len(query.Data) > 0 {
		addError(field.Name, errors, VALIDATION_ERROR_UNIQUE, &valid)
	}

	return
}

func (field *RequestField) Unique(newRule UniqueRule) *RequestField {
	var rule ValidationRule = newRule

	field.Rules = append(field.Rules, rule)

	return field
}
