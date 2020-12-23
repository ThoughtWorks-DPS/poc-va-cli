#!/bin/bash
set -e

FILE=./cover.out
THRESHOLD=90.0

if test -f "$FILE"; then
    COVERAGE=$(go tool cover -func cover.out | grep total | awk '{print substr($3, 1, length($3)-1)}')
    if [[ ( $COVERAGE < $THRESHOLD ) ]]; then
      echo "Code coverage is below the threshold. Failing pipeline..."
      exit 1
    fi
else
    echo "Error: Code Coverage report $FILE not found."
    exit 1
fi

