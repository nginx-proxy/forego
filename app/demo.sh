#!/bin/sh

NAME="$1"

sigterm() {
  echo "$NAME: got sigterm"
}

trap sigterm TERM

while true; do
  echo "$NAME: ping $$"
  sleep 60
done
