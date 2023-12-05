package main

import (
	"log"
	"os/exec"
)

func main() {
	shutDownCmd := exec.Command("sudo", "shutdown", "-h", "now")
	err := shutDownCmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
