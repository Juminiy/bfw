package utils

import (
	"strconv"
	"strings"
	"unicode"
)

// FieldNameCamelToSnakeAndAddBackticks
// UpdatedAt, updatedAt -> `updated_at`
// DeletedAt, deletedAt -> `deleted_at`
// CreatedAt, createdAt -> `created_at`
// XxxYyy, xxxYyy -> `xxx_yyy`
func FieldNameCamelToSnakeAndAddBackticks(s string) string {
	var result string
	var words []string
	l := 0
	for s != "" {
		l = strings.IndexFunc(s[1:], unicode.IsUpper) + 1
		if l <= 0 {
			break
		}
		words = append(words, s[:l])
		s = s[l:]
	}
	if s != "" {
		words = append(words, s)
	}
	for i, word := range words {
		if i > 0 {
			result += "_"
		}
		result += strings.ToLower(word)
	}
	return "`" + result + "`"
}

func InterfaceToString(inter interface{}) (string, string) {
	switch inter.(type) {
	case int:
		return "int", strconv.Itoa(inter.(int))
	case float64:
		return "float64", strconv.FormatFloat(inter.(float64), 'f', 6, 64)
	case string:
		return "string", inter.(string)
	default:
		return undefinedString, undefinedString
	}
}
