{
  description = "Kumo dependency flake wohoooo";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs = { self, nixpkgs }:
  let

    supportedSystems = [ "x86_64-linux" "aarch64-darwin" ];
    forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
    nixpkgsFor = forAllSystems (system:
      import nixpkgs { inherit system; }
    );

  in
  {
    devShells = forAllSystems (system:
      let

        pkgs = nixpkgsFor.${system};

      in {

        default = pkgs.mkShell {

          buildInputs = [
            pkgs.go_1_25
            pkgs.gopls
            pkgs.golangci-lint-langserver
            # pkgs.fish
          ];

          shellHook = ''
            # export SHELL=${pkgs.fish}/bin/fish
            # exec ${pkgs.fish}/bin/fish
          '';

        };
      }

    );
  };

}
