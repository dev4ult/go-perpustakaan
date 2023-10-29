package helpers

import (
	"strconv"
	"strings"
)

func ParseError(errMessage string) map[string]any {
	errorMap := make(map[string]any)

	// Split the error message by commas
	errorParts := strings.Split(errMessage, ",")

	// Iterate through error parts and extract key-value pairs
	for _, part := range errorParts {
		part = strings.TrimSpace(part)
		keyValue := strings.SplitN(part, "=", 2)
		if len(keyValue) == 2 {
			var value string = keyValue[1]
			if keyValue[0] == "code" {
				valueInt, _ := strconv.Atoi(value)
				errorMap[keyValue[0]] = valueInt
			} else {
				errorMap[keyValue[0]] = value
			}
		}
	}

	return errorMap
}