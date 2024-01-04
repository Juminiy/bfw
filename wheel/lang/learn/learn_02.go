package learn

import (
	"fmt"
	"strings"
)

type Animal interface {
	Speak(...string)
}

type Dog struct {
	name string
}

func (a *Dog) Speak(words ...string) {
	word := strings.Join(words, ",")
	fmt.Println(a.name, "say:", "[", word, "]")
}

type Cat struct {
	name string
}

func (a *Cat) Speak(words ...string) {
	word := strings.Join(words, "\n")
	fmt.Println(a.name, "say:", "\n", "[", word, "]")
}
