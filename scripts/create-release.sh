#!/bin/sh

set -e

VERSION=$1
BUILD_DIR=~/src/github.com/mikerybka/apps/webmachine.dev/builds/webmachine_$VERSION-1_amd64

if [ -z "$VERSION" ]; then
  echo "Usage: $0 <version>"
  exit 1
fi

mkdir -p $BUILD_DIR/DEBIAN
cat > $BUILD_DIR/DEBIAN/control <<EOF
Package: webmachine
Version: ${VERSION}
Architecture: amd64
Maintainer: Mike Rybka <merybka@gmail.com>
Description: Host static or dynamic websites with ease.
 Supports multiple domains, TLS, and several programming languages inlcuding Python, Ruby, Go, Rust and TypeScript.
 Read more at https://webmachine.dev
EOF

mkdir -p $BUILD_DIR/usr/local/bin
go build -o $BUILD_DIR/usr/local/bin/webmachine github.com/mikerybka/webmachine

dpkg-deb --build --root-owner-group $BUILD_DIR
