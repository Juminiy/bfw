package lang

import (
	"math/rand"
	"strconv"
	"strings"
	"unicode"
)

const (
	undefinedString string = ""
	plusSign        byte   = '+'
	minusSign       byte   = '-'
	timesSign       byte   = '*'
	divSign         byte   = '/'
)

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func ConcatIntToString(d string, a ...int) string {
	destStr := undefinedString
	if aLen := len(a); aLen > 0 {
		destStr += strconv.Itoa(a[0])
		for idx := 1; idx < aLen; idx++ {
			destStr += d + strconv.Itoa(a[idx])
		}
	}
	return destStr
}
func StringIsNull(str string) bool {
	return str == undefinedString
}

func Float64StringIsMinus(str string) bool {
	return str[0] == minusSign
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func String2Uint(str string) (uint, error) {
	if i, err := strconv.Atoi(str); err != nil {
		return 0, err
	} else {
		return uint(i), nil
	}
}

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
