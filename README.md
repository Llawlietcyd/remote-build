##remote build

##Build and Run

1. Generate ProtoBuf

protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     remote-build/remote-build.proto

2. Run Server

go run server/main.go

3. Run Client to test

go run client/main.go --addr localhost:50051 --filename main.c --command "gcc -o main main.c" --content "$(printf '#include <stdio.h>\nint main() {\n    printf("Hello, World!\\n");\n    return 0;\n}')"

