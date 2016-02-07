#!/bin/sh

go install github.com/murphybytes/udt.go/test/client
go install github.com/murphybytes/udt.go/test/server

server &
server_pid=$!


for teststring in "Hello World" "kdkdlleeooidididilkeadkdkdkdkdkkdldldlldldldldldldldl" "x";
do
client -tb "$teststring"
res=$?
if [ "$res" -ne "0" ]; then
  echo "Test Failed."
  break;
fi
done

kill -15 $server_pid
