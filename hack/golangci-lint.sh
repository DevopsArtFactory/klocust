#!/bin/bash
set -e -o pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
BIN=${DIR}/bin
VERSION=1.64.8

function install_linter() {
  echo "Installing GolangCI-Lint"
  GOBIN=${BIN} go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$VERSION
}

if ! [ -x "$(command -v ${BIN}/golangci-lint)" ] ; then
  install_linter
elif [[ $(${BIN}/golangci-lint --version | grep -c " $VERSION ") -eq 0 ]]
then
  echo "required golangci-lint: v$VERSION"
  echo "current version: $(${BIN}/golangci-lint --version)"
  echo "reinstalling..."
  rm -f "${BIN}/golangci-lint"
  install_linter
fi

FLAGS=""
if [[ "${CI}" == "true" ]]; then
    FLAGS="-v --print-resources-usage"
fi

GOFLAGS="${GOFLAGS:+$GOFLAGS }-mod=mod" \
${BIN}/golangci-lint run ${FLAGS} -c ${DIR}/golangci.yml \
    | awk '/out of memory/ || /Timeout exceeded/ {failed = 1}; {print}; END {exit failed}'
