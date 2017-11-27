#!/bin/sh

protoc -I /home/sbeck/go/src/github.com/gogo/protobuf/protobuf/ -I proto/ proto/*.proto --js_out=import_style=commonjs,binary:../web/gotime/src/app/proto/ --ts_out=service=true:../web/gotime/src/app/proto/ --gogofast_out=plugins=grpc:./proto


