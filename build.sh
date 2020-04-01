#!/bin/bash
echo "Kill old server"
kill -9 $(lsof -i :19999 | awk '{print $2}')

go build .

./oa