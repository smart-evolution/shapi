#!/usr/bin/sh

VERSION=$(git tag --sort version:refname | tail -n 1)

echo "Generating version="$VERSION

sed "s/%VERSION%/\"$VERSION\"/g" $1 > $2
