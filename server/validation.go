package server

import "strconv"

type RequestField struct {
	Name  string
	Rules []string
}

func (field *RequestField) Required() *RequestField {
	field.Rules = append(field.Rules, RULE_REQUIRED)

	return field
}

func (field *RequestField) Int() *RequestField {
	field.Rules = append(field.Rules, RULE_INT)

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

func Validate(rules []RequestField) (reqValues map[string]any, valid bool) {
	reqValues = map[string]any{}
	valid = true

	for _, field := range rules {
		var convertedValue any
		val := GetParam(field.Name)
		fieldValid := true

		for _, rule := range field.Rules {
			switch rule {
			case RULE_REQUIRED:
				if val == "" {
					fieldValid = false
				}
			case RULE_INT:
				value, err := strconv.Atoi(val)
				if err != nil {
					fieldValid = false
				}
				convertedValue = value
			}
		}

		if fieldValid {
			if convertedValue != nil {
				reqValues[field.Name] = convertedValue
			} else {
				reqValues[field.Name] = val
			}
		} else {
			valid = false
		}
	}

	return
}
