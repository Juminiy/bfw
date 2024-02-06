package string

import "strings"

type SQLBuilder struct {
	builder strings.Builder
}

func (b *SQLBuilder) String() string {
	return b.builder.String()
}

func (b *SQLBuilder) Reset() {
	b.builder.Reset()
}

func (b *SQLBuilder) beginQuote(q byte) {

}

func (b *SQLBuilder) endQuote(q byte) {

}
