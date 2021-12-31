{ pkgs ? import <nixpkgs> { system = builtins.currentSystem;}}:

with pkgs;

mkShell {
  buildInputs = [ bazel_4 go ];
}
