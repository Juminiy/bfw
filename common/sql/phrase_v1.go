package sql

import (
	"errors"
	"regexp"
	"strings"
)

const (
	undefinedString         string = ""
	SpacerSymbol            string = " "
	DelimiterDot            string = "."
	SpacerComma             string = ","
	TerminalSymbolSemicolon string = ";"
)

func RemoveSqlPhrasePairCharacter(phrase string, char ...rune) string {
	charLen := len(char)
	switch charLen {
	case 0:
		{
			break
		}
	case 1:
		{
			charStr := string(char)
			if strings.HasPrefix(phrase, charStr) && strings.HasSuffix(phrase, charStr) {
				return strings.ReplaceAll(phrase, charStr, undefinedString)
			}
			break
		}
	case 2:
		{
			leftChar, rightChar := string(char[0]), string(char[1])
			if strings.HasPrefix(phrase, leftChar) && strings.HasSuffix(phrase, rightChar) {
				tPhrase := strings.Replace(phrase, leftChar, undefinedString, 1)
				return strings.Replace(tPhrase, rightChar, undefinedString, 1)
			}
			break
		}
	default:
		{
			break
		}
	}

	return phrase
}

func AddSqlPhrasePairCharacter(phrase string, char ...rune) string {
	charLen := len(char)
	switch charLen {
	case 0:
		{
			break
		}
	case 1:
		{
			charStr := string(char)
			if !(strings.HasPrefix(phrase, charStr) && strings.HasSuffix(phrase, charStr)) {
				return charStr + phrase + charStr
			}
			break
		}
	case 2:
		{
			leftChar, rightChar := string(char[0]), string(char[1])
			if !(strings.HasPrefix(phrase, leftChar) && strings.HasSuffix(phrase, rightChar)) {
				return leftChar + phrase + rightChar
			}
			break
		}
	default:
		{
			break
		}
	}

	return phrase
}

func RemoveSqlPhraseBacktick(phrase string) string {
	return RemoveSqlPhrasePairCharacter(phrase, '`')
}

func AddSqlPhraseBacktick(phrase string) string {
	return AddSqlPhrasePairCharacter(phrase, '`')
}

func RemoveSqlPhraseBracket(phrase string) string {
	return RemoveSqlPhrasePairCharacter(phrase, '(', ')')
}

func AddSqlPhraseBracket(phrase string) string {
	return AddSqlPhrasePairCharacter(phrase, '(', ')')
}

func RemoveSqlPhraseQuotationMark(phrase string) string {
	return RemoveSqlPhrasePairCharacter(phrase, '"')
}

func AddSqlPhraseQuotationMark(phrase string) string {
	return AddSqlPhrasePairCharacter(phrase, '"')
}

func AddSqlPhraseApostrophe(phrase string) string {
	return AddSqlPhrasePairCharacter(phrase, '\'')
}

func AddSqlPhrasePercentSign(phrase string) string {
	return AddSqlPhrasePairCharacter(phrase, '%')
}

func AddSqlPhraseBracketDirectly(phrase string) string {
	return "(" + phrase + ")"
}

func AddSqlPhrasePercentSignDirectly(phrase string) string {
	return "%" + phrase + "%"
}

func AddSqlPhraseQuotationMarkDirectly(phrase string) string {
	return "\"" + phrase + "\""
}

// SeparateCompleteColumnNameByDot
// src data: `tbl_1`.`col_1`
// dest data: `tbl_1`, `col_1`
func SeparateCompleteColumnNameByDot(mixColumnName string) (string, string) {
	parts := strings.Split(mixColumnName, ".")
	if len(parts) == 2 {
		return parts[0], parts[1]
	} else {
		return "", mixColumnName
	}
}

func SplitSqlPhraseAndWrapCompleteBacktick(phrase string) string {
	table, column := SeparateCompleteColumnNameByDot(phrase)
	return AddSqlPhraseBacktick(RemoveSqlPhraseBacktick(table) + "." + RemoveSqlPhraseBacktick(column))
}

func ExtractSqlPhraseTableNames(phrase string) []string {
	re := regexp.MustCompile(`(?i)(?:FROM|JOIN|UPDATE|INTO)\s+([^\s,;]+)`)
	matches := re.FindAllStringSubmatch(phrase, -1)
	tableNames := make([]string, 0, len(matches))
	for _, match := range matches {
		tableName := strings.TrimSpace(match[1])
		tableNames = append(tableNames, tableName)
	}
	return tableNames
}

func ExtractSqlPhraseMainTableName(phrase string) string {
	re := regexp.MustCompile(`(?i)FROM\s+(\S+)`)
	matches := re.FindStringSubmatch(phrase)
	if len(matches) > 1 {
		return matches[1]
	}
	return undefinedString
}

func ExtractSqlPhraseAliasTableName(phrase string) string {
	re := regexp.MustCompile(`(?i)\(\s*SELECT\s+.+\s*\)\s+AS\s+(\S+)`)
	matches := re.FindStringSubmatch(phrase)
	if len(matches) > 1 {
		return matches[1]
	}
	return undefinedString
}

// JudgeSqlPhraseContainSubQuery
// GROUP BY and SubQuery
func JudgeSqlPhraseContainSubQuery(phrase string) bool {
	re := regexp.MustCompile(`(?i)\(\s*SELECT\s+.+\s*\)`)
	matches := re.FindStringSubmatch(phrase)
	return len(matches) > 0 && strings.Contains(strings.ToUpper(phrase), "GROUP BY")
}

func GetSqlPhraseAggregationAliasColumnNameMysqlDataType(phrase string) string {
	if phrase != undefinedString {
		if strings.HasSuffix(phrase, "_count") ||
			strings.HasSuffix(phrase, "_sum") ||
			strings.HasSuffix(phrase, "_max") ||
			strings.HasSuffix(phrase, "_min") {
			return "INT"
		} else if strings.HasSuffix(phrase, "_avg") {
			return "FLOAT"
		}
	}
	return undefinedString
}

func CutSqlPhraseSuffix(phrase string, cutLen int) string {
	sqlPhraseLen := len(phrase)
	if sqlPhraseLen >= cutLen {
		return phrase[:sqlPhraseLen-cutLen]
	} else {
		return phrase
	}
}

func RemoveSqlPhrasePrefixSpace(phrase string) string {
	chIdx := 0
	for idx, ch := range phrase {
		if !(ch == ' ' || ch == '\n' || ch == '\t') {
			chIdx = idx
			break
		}
	}
	return phrase[chIdx:]
}

func RemoveSqlPhraseSuffixSpace(phrase string) string {
	chIdx := len(phrase) - 1
	for ; chIdx >= 0; chIdx-- {
		if ch := phrase[chIdx]; !(ch == ' ' || ch == '\n' || ch == '\t') {
			break
		}
	}
	return phrase[:chIdx+1]
}

func ValidateDataQuerySqlPhrase(phrase string) error {
	// remove prefix space and tab
	phrase = RemoveSqlPhrasePrefixSpace(phrase)
	if !strings.HasPrefix(phrase, "SELECT") {
		return errors.New("sqlPhrase has not SELECT error")
	}
	return nil
}

// FindPrecedingOnePhrase
// "SELECT mi   " -> "SELECT", "mi"
func FindPrecedingOnePhrase(phrase string) (string, string) {
	phrase = RemoveSqlPhraseSuffixSpace(phrase)
	chIdx, found := len(phrase)-1, false
	for ; chIdx >= 0; chIdx-- {
		if ch := phrase[chIdx]; ch == ' ' || ch == '\n' || ch == '\t' {
			found = true
			break
		}
	}
	if found {
		return phrase[:chIdx+1], phrase[chIdx+1:]
	} else {
		return phrase, undefinedString
	}
}
