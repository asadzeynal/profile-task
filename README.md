# Profile-task

To run this project you need to :
1. Clone this repo
2. Run:
    ```
    make build
    make run
    ```
Grpc runs on port 8080, gateway on 8081 and swaggerui on 8082.

This project uses [buf](https://buf.build/docs/installation) (wrapper on protoc) to generate code and docs from proto. It also handles external dependencies.

http handler available at :8081/api/getprofile?inn=[xxxxxxxx]

swagger available at :8082/swaggerui