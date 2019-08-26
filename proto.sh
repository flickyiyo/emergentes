#!/bin/bash
protoc imgstream/imgstream.proto --go_out=plugins=grpc:.
protoc imgstream/imgstream.proto --python_out=./python
