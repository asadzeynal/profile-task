proto: 
	rm -rf gen/*
	buf generate proto

gateway:
	go run ./cmd/http/main.go

grpc: 
	go run ./cmd/grpc/main.go


.PHONY: proto
