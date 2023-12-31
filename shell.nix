{ pkgs ? import <nixpkgs> { } }:

with pkgs;

mkShell {
  buildInputs = [
    docker-compose-language-service
    dockfmt
    go
    gopls
    gofumpt
    golangci-lint
    golangci-lint-langserver
    golines
    delve

    gccStdenv
  ];
}
