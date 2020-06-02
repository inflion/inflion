#!/bin/bash

docker run --net host --rm \
  -v "$(pwd)/../../dbconfig.yml:/workspace/dbconfig.yml" \
  -v "$(pwd)/../../internal/store/schema:/workspace/internal/store/schema" \
  pyar6329/sql-migrate:latest up
