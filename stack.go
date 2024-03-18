package main

import "os"

type Stack string
const (
	GOLANG Stack = "Go"
	TERRAFORM Stack = "Terraform"
	UNKNOWN Stack = "Unknown"
)

func GetStack() Stack {
	if fileExists("go.mod") {
		return GOLANG
	}
	if fileExists("main.tf") {
		return TERRAFORM
	}
	return UNKNOWN
}

