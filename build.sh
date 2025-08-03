#!/bin/sh
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o course
docker build -t ccr.ccs.tencentyun.com/mryao/course:250803001 .
docker push ccr.ccs.tencentyun.com/mryao/course:250803001
rm -rf course