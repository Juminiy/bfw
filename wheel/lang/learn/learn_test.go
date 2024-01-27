package learn

import (
	"fmt"
	"reflect"
	"testing"
	"time"
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

func TestNeko(t *testing.T) {

	type nekoRes struct {
		nekoPtr *Neko
		err     error
	}
	nekoResChan := make(chan nekoRes)
	go func() {
		anySlice := []any{Neko{}, &Neko{}, []Neko{}, []*Neko{}, [2]Neko{}, [2]*Neko{}}
		for i := 0; i < len(anySlice); i++ {
			time.Sleep(1 * time.Second)
			nekoPtr, err := ReadAny2NekoPtr(anySlice[i])
			nekoResChan <- nekoRes{nekoPtr: nekoPtr, err: err}
		}
	}()
	cnt := 0
	for {
		select {
		case res := <-nekoResChan:
			cnt++
			fmt.Println(res)
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("tasks are running")
		}
		if cnt == 6 {
			break
		}
	}

}

func TestGetNone(t *testing.T) {
	a := GetNone()
	fmt.Println("original one", a)
	setValue := func(v reflect.Value) {
		if v.IsNil() {
			switch v.Kind() {
			case 1:
				v.Set(reflect.ValueOf(true))
			case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
				v.Set(reflect.ValueOf(1))
			case 13, 14:
				v.Set(reflect.ValueOf(1.25))
			case 15, 16:
				v.Set(reflect.ValueOf(1 + 2i))
			case reflect.String:
				v.Set(reflect.ValueOf("not nil"))
			default:
				if v.CanAddr() {
					v.Set(reflect.ValueOf(nil))
				}
			}
		}
	}

	type None struct {
		A int
		B string
		C complex128
	}

StructHere:
	aType := reflect.TypeOf(a)
	aValue := reflect.ValueOf(a)

	switch aType.Kind() {
	case reflect.Struct:
		for i := 0; i < aType.NumField(); i++ {
			setValue(aValue.Field(i))
			fmt.Println("field[", i, "]:", aType.Field(i).Name)
		}
	case reflect.Pointer:
		if aTypeString := aType.String(); aTypeString == "*learn.None" {
			fmt.Println("assert success")
			a = a.(*None)
			goto StructHere
		} else {
			fmt.Println("assert failed")
		}
	default:
		panic("GetNone ret unknown Type")
	}
	fmt.Println("reset to", a)
}
