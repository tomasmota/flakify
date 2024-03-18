package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/template"

	"github.com/tomasmota/flakify/stack"
)

func main() {
	err := exec.Command("nix", "--version").Run()
	if err != nil {
		log.Fatal("Nix is not installed in the system")
	}

	if !flakesEnabled() {
		log.Fatal("Flakes are not enabled")
	}

	stack, err := stack.GetStack()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Identified stack:", stack.Name())
	generateFlake(stack.GetTemplate(), "nixi") // take project name for stdin
}

func generateFlake(flakeTemplate string, projectName string) error {
	tmpl, err := template.New("flake").Parse(flakeTemplate)
	if err != nil {
		return err
	}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, struct {
		ProjectName string
	}{
		ProjectName: projectName,
	})
	if err != nil {
		return err
	}

	err = os.WriteFile("gen.nix", tpl.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}
