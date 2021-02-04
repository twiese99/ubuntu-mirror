#!/bin/bash

args=()

if [ -n "${MIRROR_TYPE}" ]; then
  args+=(--type "$MIRROR_TYPE")
fi
if [ -n "${INTERVAL}" ]; then
  args+=(--interval "$INTERVAL")
fi
if [ -n "${COUNTRY_CODE}" ]; then
  args+=(--country "$COUNTRY_CODE")
fi
if [ -n "${RSYNC_SOURCE}" ]; then
  args+=(--source "$RSYNC_SOURCE")
fi

echo ./ubuntu-mirror "${args[@]}"
./ubuntu-mirror "${args[@]}"
