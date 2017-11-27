#!/bin/sh

echo $BINDIR
protoc -I proto/ proto/*.proto --js_out=import_style=commonjs,binary:../web/gotime/src/app/proto/ --ts_out=service=true:../web/gotime/src/app/proto/ --gogofaster_out=plugins=grpc:proto/


