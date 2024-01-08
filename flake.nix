{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    senv.url = "github:luisnquin/senv";
  };

  outputs = inputs:
    with inputs;
      flake-utils.lib.eachDefaultSystem (
        system: {
          defaultPackage = pkgs.hello;

          devShells.default = let
            pkgs = import nixpkgs {
              config = {
                allowBroken = false;
                allowUnfree = true;
              };
              inherit system;
            };
          in
            pkgs.mkShell {
              inherit system;

              buildInputs = [
                senv.defaultPackage.${system}
                pkgs.golangci-lint
                pkgs.panicparse
                pkgs.go_1_21
                pkgs.hadolint
                pkgs.gofumpt
                pkgs.just
                pkgs.sqlc
                pkgs.git
                pkgs.air
              ];
            };
        }
      );
}
