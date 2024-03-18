package stack

type Terraform struct{}

func (g* Terraform) Name() string {
	return "terraform"
}

func (g *Terraform) Identify() bool {
	return fileExists("main.tf")
}

func (g *Terraform) GetTemplate() string {
	return `
{
  description = "Flake for Terraform development";

  inputs = {
    nixpkgs.url = "https://flakehub.com/f/NixOS/nixpkgs/0.1.0.tar.gz";
  };


  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system ; };
      in
      {
        devShell = pkgs.mkShellNoCC {
          packages = with pkgs; [
            terraform
            tflint
            terraform-docs
          ];
          shellHook = ''
            # Set some useful aliases
            alias tfi="terraform init"
            alias tfp="terraform plan"
            alias tfa="terraform apply"
            alias tfd="terraform destroy"
            alias tfs="terraform show"
            alias tfv="terraform validate"
          '';
        };
      }
    );
}
`
}
