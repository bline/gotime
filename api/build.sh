#!/bin/bash

bindir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
web_out="$bindir/../web/gotime/src/app/proto/"
go_out="$bindir/proto/"

protoc -I proto/ proto/*.proto \
	--js_out=import_style=commonjs,binary:$web_out \
	--ts_out=service=true:$web_out \
	--gogofaster_out=plugins=grpc:$go_out


