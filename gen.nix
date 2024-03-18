
{
  description = "Flake for nixi";

  inputs = {
    nixpkgs.url = "https://flakehub.com/f/NixOS/nixpkgs/0.1.0.tar.gz";
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
        packages.nixi = pkgs.buildGoModule {
          name = "nixi";
          src = gitignore.lib.gitignoreSource ./.;
          vendorHash = null;
        };

        packages.default = packages.nixi;

        devShell = pkgs.mkShellNoCC {
          packages = with pkgs; [
            go_1_22
            gotools
            golangci-lint
          ];
        };
      }
    );
}
