#!/bin/bash

INSTALL_PREFIX=/usr
INSTALL_DIR="$INSTALL_PREFIX/bin"
MAN_DIR="$INSTALL_PREFIX/share/man/man6"
BASH_COMPLETION_DIR="$INSTALL_PREFIX/share/bash-completion/completions"
ZSH_COMPLETION_DIR="$INSTALL_PREFIX/share/zsh/site-functions"

function output_and_create_dir() {
	echo "Installing tmatrix $1 to $2"
	mkdir -p "$2"
}

output_and_create_dir "binary" "$INSTALL_DIR"
install tmatrix "$INSTALL_DIR"
output_and_create_dir "man" "$MAN_DIR"
install -g 0 -o 0 -m 0644 tmatrix.6 "$MAN_DIR"
gzip -f "$MAN_DIR/tmatrix.6"
output_and_create_dir "bash completion script" "$BASH_COMPLETION_DIR"
install -g 0 -o 0 -m 0644 completions/tmatrix-completion.bash "$BASH_COMPLETION_DIR/tmatrix"
output_and_create_dir "zsh completion script" "$ZSH_COMPLETION_DIR"
install -g 0 -o 0 -m 0644 completions/tmatrix-completion.zsh "$ZSH_COMPLETION_DIR/_tmatrix"
