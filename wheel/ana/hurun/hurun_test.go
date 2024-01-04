package hurun

import (
	"fmt"
	"testing"
)

func TestMakeHurunURL(t *testing.T) {
	total := GetUniCornEnterpriseTotalCount()
	uniCorn := GetUniCornEnterpriseParse2JSONFile("hurun_2023", "", total)
	fmt.Println(uniCorn.GetCompanyNameList(10))
	fmt.Println(uniCorn.GetFounderNameList(10))
}
