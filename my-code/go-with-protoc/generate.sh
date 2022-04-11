#! /bin/bash

protoc -I src/ --go_out=src/ src/protofiles/*.proto