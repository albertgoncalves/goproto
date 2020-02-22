with import <nixpkgs> {};
mkShell {
    buildInputs = [
        go
        golint
        shellcheck
    ];
    shellHook = ''
        . .shellhook
    '';
}
