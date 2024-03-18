package stack

import "os"

type Stack interface {
	Identify() bool 
	GetTemplate() string
}

func fileExists(fileName string) bool {
	if _, err := os.Stat(fileName); err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
