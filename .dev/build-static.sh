#!/bin/sh

go build --ldflags "-X 'main.Version=$(git describe --tags)' -linkmode external -extldflags \"-static\" -s -w" -o rofixec .
