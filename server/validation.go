package server

import (
	"fmt"
	"strconv"
)

type RequestField struct {
	Name string
	Rule ValidationRule
}

type ValidationRule struct {
	Required bool
	Length   int
	Type     string
	Values   []any
}

func (field *RequestField) Required() *RequestField {
	field.Rule.Required = true

	return field
}

func (field *RequestField) Int() *RequestField {
	field.Rule.Type = TYPE_INT

	return field
}

func (field *RequestField) String() *RequestField {
	field.Rule.Type = TYPE_STRING

	return field
}

func (field *RequestField) Sometimes() *RequestField {
	field.Rule.Required = false

	return field
}

func (field *RequestField) In(values []interface{}) *RequestField {
	field.Rule.Values = values

	return field
}

func (field *RequestField) SetName(fieldName string) *RequestField {
	field.Name = fieldName

	return field
}

func Rule(fieldName string) (reqField *RequestField) {
	reqField = new(RequestField).SetName(fieldName)
	reqField.Required()

	return
}

func Validate(reqFields []RequestField) (reqValues map[string]any, valid bool, fieldErrors map[string][]string) {
	reqValues = make(map[string]any)
	valid = true

	for _, field := range reqFields {
		var convertedValue interface{}
		fieldValid := true
		value := GetParam(field.Name)

		if value == "" {
			if field.Rule.Required {
				addError(field.Name, &fieldErrors, fmt.Sprintf(VALIDATION_ERROR_REQUIRED, field.Name), &fieldValid)
			} else {
				continue
			}
		}

		if field.Rule.Type != "" {
			switch field.Rule.Type {
			case TYPE_INT:
				intValue, error := strconv.Atoi(value)
				if error == nil {
					convertedValue = intValue
				} else {
					addError(field.Name, &fieldErrors, fmt.Sprintf(VALIDATION_ERROR_INCORRECT_TYPE, field.Name, TYPE_INT), &fieldValid)
				}
			case TYPE_STRING:
				convertedValue = value
			}
		}

		// valueType := reflect.TypeOf(convertedValue).String()

		// fmt.Println(valueType)
		// var isInSlice bool
		// switch val := convertedValue.(type) {
		// case int:
		// 	isInSlice = helper.InSlice(field.Rule.Values, val)
		// }

		if fieldValid {
			if convertedValue != nil {
				reqValues[field.Name] = convertedValue
			} else if value != "" {
				reqValues[field.Name] = value
			}
		} else {
			valid = false
		}
	}

	if fieldErrors != nil {
		valid = false
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
