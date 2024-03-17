package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/template"
)

func main() {
	err := exec.Command("nix", "--version").Run()
	if err != nil {
		log.Fatal("Nix is not installed in the system")
	}

	if !flakesEnabled() {
		log.Fatal("Flakes are not enabled")
	}

	stack:= GetStack()
	fmt.Println("Stack in current directory:", stack)
	generateGoFlake("nixi")
}

const goFlakeTemplate = `
{
  description = "Flake for {{ .ProjectName }}";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    gitignore = {
      url = "github:hercules-ci/gitignore.nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs = { self, nixpkgs, flake-utils, gitignore }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system ; };
      in
      rec {
        packages.{{ .ProjectName }} = pkgs.buildGoModule {
          name = "{{ .ProjectName }}";
          src = gitignore.lib.gitignoreSource ./.;
          vendorHash = null;
        };

        packages.default = packages.{{ .ProjectName }};

        devShell = pkgs.mkShellNoCC {
          packages = with pkgs; [
            go_1_22
            gotools
            gopls
            golangci-lint
          ];
        };
      }
    );
}
`

func generateGoFlake(projectName string) error {
	tmpl, err := template.New("flake").Parse(goFlakeTemplate)
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

	err = os.WriteFile("flake.nix", tpl.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}
