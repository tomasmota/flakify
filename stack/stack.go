package stack

import (
	"errors"
	"os"
)

type Stack interface {
	Name() string
	Identify() bool
	GetTemplate() string
}

var stacks = []Stack{
	&Golang{},
	&Terraform{},
	&Poetry{},
}

func GetStack() (Stack, error) {
	for _, stack := range stacks {
		if stack.Identify() {
			return stack, nil
		}
	}
	return nil, errors.New("no stack identified for current project")
}

func fileExists(fileName string) bool {
	if _, err := os.Stat(fileName); err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
