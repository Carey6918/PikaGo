#!/usr/bin/env bash
set -x

CUR_DIR=$(cd $(dirname $0); pwd)
CUR_DATE=$(date +%Y-%m-%d)
export CONF_DIR=$CUR_DIR/conf

go run main.go
