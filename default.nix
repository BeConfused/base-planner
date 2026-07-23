{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  buildInputs = with pkgs; [
    go_1_26
    git
    nodejs_26
    pnpm
  ];
  GOPATH="${toString ./.}/.nix-workspace";
  PLAN_FILE="test-configs/test-plan.yaml";
  RECIPE_FILE="test-configs/recipes.yaml";
  MATERIALS_FILE="test-confis/materials.yaml";
}
