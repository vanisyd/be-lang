package validator

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"web/helper"
)

// Base rule
type ValidationRule interface {
	Check(field *RequestField, errors *map[string][]string) (valid bool)
}

// Required
type RequiredRule struct{}

func (rule RequiredRule) Check(field *RequestField, errors *map[string][]string) (valid bool) {
	valid = true

	if field.Value == "" {
		addError(field.Name, errors, fmt.Sprintf(VALIDATION_ERROR_REQUIRED, field.Name), &valid)
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
			addError(field.Name, errors, fmt.Sprintf(VALIDATION_ERROR_INCORRECT_TYPE, field.Name, TYPE_INT), &valid)
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

	valid = helper.InSlice(rule.Values, field.ConvertedValue.(T))

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
