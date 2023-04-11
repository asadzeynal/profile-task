#!/bin/sh

/app/grpc &

/app/gateway &
  
./swagger &

wait -n
  
exit $?
