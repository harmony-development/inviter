#!/usr/bin/env bash
mkdir -p "gen"

cp -R ./protocol ./protocol-build-tmp

find ./protocol-build-tmp \( -type d -name .git -prune \) -o -type f -print0 | xargs -0 sed -i 's/github\.com\/harmony-development\/legato/github\.com\/harmony-development\/inviter/g'

for dir in $(find "protocol-build-tmp" -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq); do
  echo $(find "${dir}" -name '*.proto')
  protoc --experimental_allow_proto3_optional \
  --proto_path=protocol-build-tmp \
  --proto_path=$(go env GOPATH)/src/github.com/envoyproxy/protoc-gen-validate \
  --go_out=./gen \
  --go_opt="plugins=grpc" \
  --validate_out="lang=go:gen" \
  $(find "${dir}" -name '*.proto')
done

rsync -a -v gen/github.com/harmony-development/inviter/gen/ ./gen
rm -r ./protocol-build-tmp