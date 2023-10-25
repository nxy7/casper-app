{
  description = "Project starter";
  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts";
    nix2container.url = "github:nlewo/nix2container";
    rust-overlay = {
      url = "github:oxalica/rust-overlay";
      inputs = { nixpkgs.follows = "nixpkgs"; };
    };

  };

  outputs = { flake-parts, nixpkgs, ... }@inputs:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = [ "x86_64-linux" ];
      perSystem = { config, system, ... }:
        let
          overlays = [ (import inputs.rust-overlay) ];
          pkgs = import nixpkgs { inherit system overlays; };
        in {
          devShells.default = pkgs.mkShell {
            packages = with pkgs; [
              (rust-bin.nightly."2023-06-15".default.override {
                extensions = [ "rust-src" ];
                targets = [ "wasm32-unknown-unknown" ];
              })
              bun
              go
              cmake
              pkg-config
              openssl.dev
            ];
          };
        };
    };
}

