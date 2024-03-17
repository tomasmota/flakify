package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	err := exec.Command("nix", "--version").Run()
	if err != nil {
		log.Fatal("Nix is not installed in the system")
	}

	if !flakesEnabled() {
		log.Fatal("Flakes are not enabled")
	}

	stack:= getStack()

	fmt.Println(stack)
}

type Stack string
const (
	GOLANG Stack = "GOLANG"
	TERRAFORM Stack = "TERRAFORM"
	UNKNOWN Stack = "UNKNOWN"
)

func getStack() Stack {
	if fileExists("go.mod") {
		return GOLANG
	}
	if fileExists("main.tf") {
		return TERRAFORM
	}
	return UNKNOWN
}

func fileExists(fileName string) bool {
	if _, err := os.Stat(fileName); err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
