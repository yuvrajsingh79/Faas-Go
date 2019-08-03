#!/bin/bash  

cd grpc/grpc-server
kubectl apply -f service.yaml
sleep 20
kubectl apply -f deployment.yaml
sleep 10

cd ../grpc-client
kubectl apply -f service.yaml
sleep 20
kubectl apply -f deployment.yaml
sleep 10
