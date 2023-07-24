package helper

import (
	"fmt"
	"regexp"
	"strings"
)

func Plural(word string) (result string) {
	result = word
	if strings.ToLower(result[len(result)-1:]) == "s" {
		result += "es"
	} else {
		result += "s"
	}

	return
}

func Singular(word string) (result string) {
	result = word
	tmpStr := strings.ToLower(result)
	fmt.Println(tmpStr[len(tmpStr)-1:])
	if tmpStr[len(tmpStr)-3:] == "ses" {
		result = result[len(result)-2:]
	} else if tmpStr[len(tmpStr)-1:] == "s" {
		result = result[:len(result)-1]
	}

	return
}

func SnakeCase(word string) (result string) {
	pattern := regexp.MustCompile(`[A-Z][^A-Z]*`)
	parts := pattern.FindAllString(word, -1)
	for i, part := range parts {
		if i > 0 {
			result += "_"
		}
		result += strings.ToLower(part)
	}

	return result
}

func CamelCase(word string) (result string) {
	parts := strings.Split(word, "_")
	for _, part := range parts {
		result += strings.ToUpper(part[:1]) + part[1:]
	}

	return result
}
