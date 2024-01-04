package learn

import (
	"testing"
)

func TestDefinedStruct_Dec(t *testing.T) {
	d := DefinedStruct{}
	(d).Dec().Print()
}

func TestCat_Speak(t *testing.T) {
	var (
		dog, cat Animal
	)
	dog = &Dog{name: "json"}
	cat = &Cat{name: "map"}

	dog.Speak("wang", "wang", "wang")
	cat.Speak("miao", "miao", "miao")
}
