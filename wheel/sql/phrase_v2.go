package sql

import (
	"bfw/wheel/lang"
	"strings"
)

func StringToSQLPhrase(phrase string) *Phrase {
	return &Phrase{phrase}
}

// Phrase redefined the wheel string by wheel SQLPhrase
// 1. support chained options
// 2. wheel Converter string <---> SQLPhrase
type Phrase struct {
	phrase string
}

func (phrase *Phrase) ToString() string {
	return phrase.phrase
}

func (phrase *Phrase) Empty() bool {
	phrase.phrase = RemoveSqlPhrasePrefixSpace(phrase.phrase)
	phrase.phrase = RemoveSqlPhraseSuffixSpace(phrase.phrase)
	return len(phrase.phrase) == 0
}

func (phrase *Phrase) Self() *Phrase {
	return phrase
}

func (phrase *Phrase) Add(addend *Phrase) *Phrase {
	phrase.phrase = phrase.ToString() + addend.ToString()
	return phrase
}

func (phrase *Phrase) AddNextPhrase(addend *Phrase) *Phrase {
	phrase.phrase += SpacerSymbol + addend.ToString()
	return phrase
}

func (phrase *Phrase) AddNextPhrases(addend ...*Phrase) *Phrase {
	for _, add := range addend {
		phrase.phrase += SpacerSymbol + add.ToString()
	}
	return phrase
}

func (phrase *Phrase) AddNextString(addendStr string) *Phrase {
	phrase.phrase += SpacerSymbol + addendStr
	return phrase
}

func (phrase *Phrase) AddNextStrings(addendStr ...string) *Phrase {
	for _, str := range addendStr {
		phrase.phrase += SpacerSymbol + str
	}
	return phrase
}

// RemoveBackticks
//
// `tbl_1` -> tbl_1
//
//	tbl_1  -> tbl_1
func (phrase *Phrase) RemoveBackticks() *Phrase {
	phrase.phrase = RemoveSqlPhraseBacktick(phrase.phrase)
	return phrase
}

// WrapBackticks
//
//	tbl_1  -> `tbl_1`
//
// `tbl_1` -> `tbl_1`
func (phrase *Phrase) WrapBackticks() *Phrase {
	phrase.phrase = AddSqlPhraseBacktick(phrase.phrase)
	return phrase
}

// RemoveBrackets
// ( )
func (phrase *Phrase) RemoveBrackets() *Phrase {
	phrase.phrase = RemoveSqlPhraseBracket(phrase.phrase)
	return phrase
}

func (phrase *Phrase) WrapBrackets() *Phrase {
	phrase.phrase = AddSqlPhraseBracket(phrase.phrase)
	return phrase
}

// RemoveQuotationMarks
// " "
func (phrase *Phrase) RemoveQuotationMarks() *Phrase {
	phrase.phrase = RemoveSqlPhraseQuotationMark(phrase.phrase)
	return phrase
}

func (phrase *Phrase) WrapQuotationMarks() *Phrase {
	phrase.phrase = AddSqlPhraseQuotationMark(phrase.phrase)
	return phrase
}

// 无论是否重复，直接添加

func (phrase *Phrase) WrapBracketsDirectly() *Phrase {
	phrase.phrase = AddSqlPhraseBracketDirectly(phrase.phrase)
	return phrase
}

func (phrase *Phrase) WrapPercentSignsDirectly() *Phrase {
	phrase.phrase = AddSqlPhrasePercentSignDirectly(phrase.phrase)
	return phrase
}

func (phrase *Phrase) WrapQuotationMarksDirectly() *Phrase {
	phrase.phrase = AddSqlPhraseQuotationMarkDirectly(phrase.phrase)
	return phrase
}

// 前一个短语直接加符号

// WrapPrecedingOnePhraseBackticks
// SELECT -> SELECT
// SELECT wiwi -> SELECT `wiwi`
// SELECT `wiwi`, kaga -> SELECT `wiwi`, `kaga`
func (phrase *Phrase) WrapPrecedingOnePhraseBackticks() *Phrase {
	originalPhrase, precedePhrase := FindPrecedingOnePhrase(phrase.phrase)
	if precedePhrase != undefinedString {
		phrase.phrase = originalPhrase + AddSqlPhraseBacktick(precedePhrase)
	}
	return phrase
}

// ( )

func (phrase *Phrase) WrapPrecedingOnePhraseBrackets() *Phrase {
	originalPhrase, precedePhrase := FindPrecedingOnePhrase(phrase.phrase)
	if precedePhrase != undefinedString {
		phrase.phrase = originalPhrase + AddSqlPhraseBracket(precedePhrase)
	}
	return phrase
}

// ' '

func (phrase *Phrase) WrapPrecedingOnePhraseApostrophe() *Phrase {
	originalPhrase, precedePhrase := FindPrecedingOnePhrase(phrase.phrase)
	if precedePhrase != undefinedString {
		phrase.phrase = originalPhrase + AddSqlPhraseApostrophe(precedePhrase)
	}
	return phrase
}

// % %

func (phrase *Phrase) WrapPrecedingOnePercentSignsPhrase() *Phrase {
	originalPhrase, precedePhrase := FindPrecedingOnePhrase(phrase.phrase)
	if precedePhrase != undefinedString {
		phrase.phrase = originalPhrase + AddSqlPhrasePercentSign(precedePhrase)
	}
	return phrase
}

// " "

func (phrase *Phrase) WrapPrecedingOnePhraseQuotationMarks() *Phrase {
	originalPhrase, precedePhrase := FindPrecedingOnePhrase(phrase.phrase)
	if precedePhrase != undefinedString {
		phrase.phrase = originalPhrase + AddSqlPhraseQuotationMark(precedePhrase)
	}
	return phrase
}

// 切断末尾长度

func (phrase *Phrase) RemoveTrailingSuffix(sufLen int) *Phrase {
	phrase.phrase = CutSqlPhraseSuffix(phrase.phrase, sufLen)
	return phrase
}

// 获取完整语句所有表名

func (phrase *Phrase) ExtractCompletedPhraseTableNames() []string {
	return ExtractSqlPhraseTableNames(phrase.phrase)
}

func (phrase *Phrase) AddCharacter(char rune) *Phrase {
	phrase.phrase += string(char)
	return phrase
}

func (phrase *Phrase) AddCharacters(char ...rune) *Phrase {
	if lang.ValidateInterfaceArrayOrSlice(char) {
		for _, ch := range char {
			phrase.AddCharacter(ch)
		}
	}
	return phrase
}

func (phrase *Phrase) RemovePrefixAllSpacerSymbol() *Phrase {
	phrase.phrase = RemoveSqlPhrasePrefixSpace(phrase.phrase)
	return phrase
}

func (phrase *Phrase) RemoveSuffixAllSpacerSymbol() *Phrase {
	phrase.phrase = RemoveSqlPhraseSuffixSpace(phrase.phrase)
	return phrase
}

func (phrase *Phrase) AddSuffixOneSpacerSymbol() *Phrase {
	return phrase.Add(&Phrase{SpacerSymbol})
}

func (phrase *Phrase) RemoveSpacerComma() *Phrase {
	phrase.RemoveSuffixAllSpacerSymbol()
	if strings.HasSuffix(phrase.phrase, SpacerComma) {
		phrase.RemoveTrailingSuffix(1)
	}
	return phrase
}

func (phrase *Phrase) AddSpacerComma() *Phrase {
	phrase.RemoveSuffixAllSpacerSymbol()
	if !strings.HasSuffix(phrase.phrase, SpacerComma) {
		phrase.phrase += SpacerComma
	}
	return phrase
}

func (phrase *Phrase) RemoveEndSemicolon() *Phrase {
	phrase.RemoveSuffixAllSpacerSymbol()
	if strings.HasSuffix(phrase.phrase, TerminalSymbolSemicolon) {
		phrase.RemoveTrailingSuffix(1)
	}
	return phrase
}

func (phrase *Phrase) AddEndSemicolon() *Phrase {
	phrase.RemoveSuffixAllSpacerSymbol()
	if !strings.HasSuffix(phrase.phrase, TerminalSymbolSemicolon) {
		phrase.phrase += TerminalSymbolSemicolon
	}
	return phrase
}

func (phrase *Phrase) ValidateCompletedQueryPhrase() error {
	return ValidateDataQuerySqlPhrase(phrase.ToString())
}

// 分离表名和列名

func (phrase *Phrase) SplitCompletedColumnNameByDot() (*Phrase, *Phrase) {
	pairParts := strings.Split(phrase.phrase, DelimiterDot)
	if len(pairParts) != 2 {
		return nil, phrase
	}
	return StringToSQLPhrase(pairParts[0]), StringToSQLPhrase(pairParts[1])
}

// RemakeCompletedColumnName
// `tbl_1`.`col_1` -> `tbl_1.col_1`
func (phrase *Phrase) RemakeCompletedColumnName() *Phrase {
	tableName, columnName := phrase.SplitCompletedColumnNameByDot()
	tableName, columnName = tableName.RemoveBackticks(), columnName.RemoveBackticks()
	if tableName != nil {
		return tableName.Add(&Phrase{DelimiterDot}).Add(columnName).WrapBackticks()
	}
	return columnName
}
