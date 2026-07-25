{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  buildInputs = with pkgs; [
    go_1_26
    git
    gitleaks
    delve
    golangci-lint
    gosec
  ];
  GOPATH="${toString ./.}/.nix-workspace";
  GOFLAGS="-o=${toString ./.}/build/";
  PLAN_FILE="${toString ./.}/test-configs/test-plan.yaml";
  CONFIG_FILE="test-configs/config.yaml";
}
