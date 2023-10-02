package utils

import (
	"strings"
)

func DeserializeStringArray(rawSql []uint8) []string {
	serialized := string(rawSql)

	if serialized == "" || serialized == "nil" {
		return nil
	}

	serialized = strings.Trim(serialized, "{}")

	elements := strings.Split(serialized, ",")

	return elements
}

func SerializeStringArray(array []string) any {
	if len(array) == 0 {
		return nil
	}

	var serialized string
	for i, item := range array {
		if i == 0 {
			serialized = "{" + item
		} else {
			serialized = serialized + "," + item
		}

		if i == len(array)-1 {
			serialized = serialized + "}"
		}
	}
	return serialized
}
