package main

import (
	"bfw/cmd/web"
)

func runWebApi() {
	web.ServeRun()
}

func main() {
	runWebApi()
}
