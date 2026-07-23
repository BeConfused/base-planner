{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  buildInputs = with pkgs; [
    go_1_26
    git
    nodejs_26
    pnpm
  ];
  GOPATH="${toString ./.}/.nix-workspace";
  GOFLAGS="-o=${toString ./.}/build/";
  PLAN_FILE="test-configs/test-plan.yaml";
  CONFIG_FILE="test-configs/config.yaml";
}
