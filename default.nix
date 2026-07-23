{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  buildInputs = with pkgs; [
    go_1_26
    git
    nodejs_26
    pnpm
  ];
  GOCMD="cmd/";
  GOPATH="${toString ./.}/.nix-workspace";
  GOFLAGS="-o=${toString ./.}/build/";
  PLAN_FILE="test-configs/test-plan.yaml";
  RECIPE_FILE="test-configs/recipes.yaml";
  MATERIALS_FILE="test-configs/materials.yaml";
}
