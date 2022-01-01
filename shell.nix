{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  name = "inviter-shell";
  buildInputs = with pkgs; [ go ];
}
