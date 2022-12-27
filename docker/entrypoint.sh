#! /usr/bin/env bash

set -e

exec "./lipsumgo-$1" "${@:2}"
