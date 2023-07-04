package validator

import "web/server/http/request"

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
		field.Value = request.GetParam(field.Name)

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
