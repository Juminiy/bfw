package main

import (
	"bfw/cmd/web"
	"bfw/common/lang"
	"fmt"
)

func runWebApi() {
	web.ServeRun()
}

func runRTest() {
	var (
		str1 string
		str2 string
	)
	_, err := fmt.Scan(&str1, &str2)
	if err != nil {
		panic(err)
	}
	str1BigNum, str2BigNum := lang.ConstructBigNum(str1), lang.ConstructBigNum(str2)
	str1BigNum.Sub(str2BigNum).Display(false, false)
}

func main() {
	runRTest()
}
