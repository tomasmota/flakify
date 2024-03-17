{
  description = "Generates flake.nix based on directory contents";

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
        packages.flakify = pkgs.buildGoModule {
          name = "flakify";
          src = gitignore.lib.gitignoreSource ./.;
          vendorHash = null;
        };

        packages.default = packages.flakify;

        devShell = pkgs.mkShellNoCC {
          packages = with pkgs; [
            go_1_21 
            gotools
            gopls
            golangci-lint
          ];
        };
      }
    );
}
