#!/bin/bash

_base="$(dirname "$0")"
_base="$(realpath "$_base")"

main() {
  source "$_base/env"
  go run "$_base/main.go"
}

main
