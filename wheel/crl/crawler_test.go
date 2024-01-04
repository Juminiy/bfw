package crl

import (
	"fmt"
	"testing"
)

func TestGetLinksByExplicitURL(t *testing.T) {
	links, err := GetLinksByExplicitURL("https://baidu.com")
	if err != nil {
		panic(err)
	}
	for _, link := range links {
		fmt.Println(link)
	}
}
