package string

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

const (
	letters    = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	lettersLen = 52
)

func randString(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = letters[rand.Intn(lettersLen)]
	}
	return string(b)
}

func Dup(b *testing.B, fn func(string, int) string) {
	src := randString(1 << 5)
	for i := 0; i < b.N; i++ {
		fn(src, 1<<13)
	}
}

func dupStringV1(src string, n int) (dest string) {
	for i := 0; i < n; i++ {
		dest += src
	}
	return
}

func dupStringV2(src string, n int) (dest string) {
	for i := 0; i < n; i++ {
		dest = fmt.Sprintf("%s%s", dest, src)
	}
	return
}

func dupStringV3(src string, n int) string {
	bd := strings.Builder{}
	for i := 0; i < n; i++ {
		bd.WriteString(src)
	}
	return bd.String()
}

func dupStringV4(src string, n int) string {
	buf := bytes.Buffer{}
	for i := 0; i < n; i++ {
		buf.WriteString(src)
	}
	return buf.String()
}

func dupStringV5(src string, n int) string {
	b := make([]byte, 0)
	for i := 0; i < n; i++ {
		b = append(b, src...)
	}
	return string(b)
}

func dupStringV6(src string, n int) string {
	b := make([]byte, 0, len(src)*n)
	for i := 0; i < n; i++ {
		b = append(b, src...)
	}
	return string(b)
}

func dupStringV7(src string, n int) string {
	bd := strings.Builder{}
	bd.Grow(len(src) * n)
	for i := 0; i < n; i++ {
		bd.WriteString(src)
	}
	return bd.String()
}

func dupStringV8(src string, n int) string {
	buf := bytes.Buffer{}
	buf.Grow(len(src) * n)
	for i := 0; i < n; i++ {
		buf.WriteString(src)
	}
	return buf.String()
}
