package main

import (
	"encoding/json"
	"log"
	"os/exec"
	"slices"
)

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
