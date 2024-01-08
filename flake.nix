{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    senv.url = "github:luisnquin/senv";
  };

  outputs = inputs:
    with inputs; let
      system = "x86_64-linux";

      pkgs = import nixpkgs {
        config = {
          allowBroken = false;
          allowUnfree = true;
        };
        inherit system;
      };
    in
      flake-utils.lib.eachDefaultSystem (
        system: {
          defaultPackage = pkgs.hello;

          devShells.default = let
            otherPkgs = [senv.defaultPackage.${system}];
          in
            pkgs.mkShell {
              buildInputs = with pkgs;
                [
                  golangci-lint
                  panicparse
                  go_1_21
                  hadolint
                  gofumpt
                  just
                  sqlc
                  git
                  air
                ]
                ++ otherPkgs;
            };
        }
      );
}
