{ pkgs ? import <nixpkgs-unstable> {} }:
  pkgs.mkShell {
    nativeBuildInputs = (with pkgs; [
      go
    ]);
}
