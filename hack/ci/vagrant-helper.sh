#!/usr/bin/env bash
set -o errexit -o nounset -o pipefail

SSH_CONFIG=".vagrant/ssh-config"
if [ ! -f "$SSH_CONFIG" ]; then
  vagrant ssh-config > "$SSH_CONFIG"
fi

exec ssh -F "$SSH_CONFIG" default sudo "$@"
