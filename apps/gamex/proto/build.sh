#!/bin/bash
# first gen msg
protoc --go_out=./ *.proto
# del msg
rm -r msg
# gen protobuf
protoc --go_out=./ *.proto
