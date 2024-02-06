package run

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func RunCompile() {
	goSrcFile := `
		package main
		import "fmt"
		func main(){
			ei := 3
			for i:=0;i<ei;i++{
				fmt.Println(i)
			}
		}
	`
	fSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fSet, "", goSrcFile, 0)
	if err != nil {
		panic(err)
	}
	err = ast.Print(fSet, astFile)
	if err != nil {
		panic(err)
	}
}
