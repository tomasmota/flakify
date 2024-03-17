package main

import (
	"encoding/json"
	"log"
	"os/exec"
	"slices"
)

func main() {
	err := exec.Command("nix", "--version").Run()
	if err != nil {
		log.Fatal("Nix is not installed in the system")
	}

	if !flakesEnabled() {
		log.Fatal("Flakes are not enabled")
	}
}

type NixConfig struct {
	ExperimentalFeatures ExperimentalFeatures `json:"experimental-features"`
}

type ExperimentalFeatures struct {
	Values []string `json:"value"`
}

func flakesEnabled() bool {
	out, err := exec.Command("nix", "show-config", "--json").Output()
	if err != nil {
		log.Fatal("ERROR")
	}

	config := &NixConfig{}
	json.Unmarshal(out, config)

	return slices.Contains(config.ExperimentalFeatures.Values, "flakes")
}
