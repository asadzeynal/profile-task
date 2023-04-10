proto: 
	rm -f pb/*.go
	protoc --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	main.proto

.PHONY: proto
