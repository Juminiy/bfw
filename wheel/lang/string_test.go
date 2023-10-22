package lang

import (
	"fmt"
	"testing"
)

func TestTruncateStringPrefixSuffixSpace(t *testing.T) {
	fmt.Println(len(TruncateStringPrefixSuffixSpace("abcde")))
	fmt.Println(len(TruncateStringPrefixSuffixSpace("  abcde")))
	fmt.Println(len(TruncateStringPrefixSuffixSpace("abcde   ")))
	fmt.Println(len(TruncateStringPrefixSuffixSpace("   ")))
	fmt.Println(len(TruncateStringPrefixSuffixSpace("  ")))
}
