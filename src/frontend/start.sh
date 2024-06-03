#!/bin/bash

_base="$(dirname "$0")"
_base="$(realpath "$_base")"

main() {
  python -m http.server -d "$_base" -b "$FLYEASY_HOST" "$FLYEASY_FRONTEND_PORT"
}

main
