package main

import "fmt"

func main() {
	var (
		s *string
	)
	fmt.Println(*s + *s)
}
