package server

type RequestField struct {
	Name  string
	Rules []string
}

func (field *RequestField) Required() *RequestField {
	field.Rules = append(field.Rules, RULE_REQUIRED)

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

func Validate(rules []RequestField) (reqValues map[string]string, valid bool) {
	reqValues = map[string]string{}
	valid = true

	for _, field := range rules {
		val := GetParam(field.Name)
		fieldValid := true

		for _, rule := range field.Rules {
			switch rule {
			case RULE_REQUIRED:
				if val == "" {
					fieldValid = false
				}
			}
		}

		if fieldValid {
			reqValues[field.Name] = val
		} else {
			valid = false
		}
	}

	return
}
