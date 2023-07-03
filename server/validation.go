package server

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"web/helper"
)

type RequestField struct {
	Name           string
	Rules          []ValidationRule
	Value          string
	ConvertedValue interface{}
}

type RequiredRule struct{}
type SometimesRule struct{}
type TypeRule struct {
	FieldType string
}
type InRule[T string | int] struct {
	Values []T
}

type ValidationRule interface {
	Check(field *RequestField, errors *map[string][]string) (valid bool)
}

func (rule RequiredRule) Check(field *RequestField, errors *map[string][]string) (valid bool) {
	valid = true

	if field.Value == "" {
		addError(field.Name, errors, fmt.Sprintf(VALIDATION_ERROR_REQUIRED, field.Name), &valid)
	}

	return
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

func (rule InRule[T]) Check(field *RequestField, errors *map[string][]string) (valid bool) {
	valid = true
	if field.Value == "" {
		return
	}

	valid = helper.InSlice(rule.Values, field.ConvertedValue.(T))

	return
}

func (field *RequestField) Required() *RequestField {
	var rule ValidationRule = RequiredRule{}
	field.Rules = append(field.Rules, rule)

	return field
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

func (field *RequestField) SetName(fieldName string) *RequestField {
	field.Name = fieldName

	return field
}

func Rule(fieldName string) (reqField *RequestField) {
	reqField = new(RequestField).SetName(fieldName)

	return
}

func Validate(reqFields []RequestField) (reqValues map[string]any, valid bool, fieldErrors map[string][]string) {
	reqValues = make(map[string]any)
	valid = true

	for _, field := range reqFields {
		fieldValid := true
		field.Value = GetParam(field.Name)

		for _, rule := range field.Rules {
			fieldValid = rule.Check(&field, &fieldErrors)
		}

		if fieldValid {
			if field.ConvertedValue != nil {
				reqValues[field.Name] = field.ConvertedValue
			} else if field.Value != "" {
				reqValues[field.Name] = field.Value
			}
		} else {
			valid = false
		}
	}

	return
}

func addError(fieldName string, errors *map[string][]string, errorText string, fieldValid *bool) {
	if (*errors) == nil {
		(*errors) = map[string][]string{}
	}
	fieldSlice, ok := (*errors)[fieldName]
	if !ok {
		(*errors)[fieldName] = []string{}
		fieldSlice = (*errors)[fieldName]
	}

	fieldSlice = append(fieldSlice, errorText)
	(*errors)[fieldName] = fieldSlice
	(*fieldValid) = false
}
