package stack

type Golang struct{}

func (g *Golang) Identify() bool {
	// TODO: build this out
	return fileExists("go.mod")
}

func GetTemplate(projectName string) string {
	return `
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
}
