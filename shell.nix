{ pkgs ? import <nixpkgs> { } }:

with pkgs;

mkShell {
  buildInputs = [
    docker-compose-language-service
    dockfmt
    go
    gofumpt
    golangci-lint
    golangci-lint-langserver
    golines
  ];
}
