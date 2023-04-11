proto: 
	rm -rf gen/*
	buf generate proto

gateway:
	go run ./cmd/http/main.go

grpc: 
	go run ./cmd/grpc/main.go

swagger:
	go run ./cmd/swagger/main.go

build:
	docker build . -t profiletask:latest

run:
	docker run -p 8080:8080 -p 8081:8081 -p 8082:8082 profiletask 

.PHONY: proto gateway grpc swagger build run 
