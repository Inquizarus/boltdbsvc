#!/usr/bin/env bash

set -euo pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )/../"
OUT_D=${OUT_D:-${DIR}/builds}
mkdir -p "$OUT_D"

# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

re='^[0-9]+([.][0-9]+)?$'
if ! [[ $GOARM =~ $re ]] ; then
  (cd cmd/golbag && GOOS=${GOOS:-linux} GOARCH=${GOARCH:-amd64} CGO_ENABLED=0 go build -o "${OUT_D}/golbag_${GOOS}_${GOARCH}")
else
 echo "build for arm"
 (cd cmd/golbag && GOOS=${GOOS:-linux} GOARCH=${GOARCH:-arm} GOARM=${GOARM:-6} CGO_ENABLED=0 go build -o "${OUT_D}/golbag_${GOOS}_${GOARCH}")
fi

