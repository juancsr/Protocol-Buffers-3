#! /bin/bash

rm -rf src/pb
protoc -I src/ --go_out=src/ src/protofiles/*.proto