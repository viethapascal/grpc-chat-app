# How to
## Dependencies
```shell
#Install grpc-web-gateway
go install github.com/grpc-ecosystem/protoc-gen-grpc-gateway-ts
# gen ts proto
protoc --grpc-gateway-ts_out=. proto/service.proto
protoc -I proto --grpc-gateway-ts_out=ts_import_roots=$(pwd),ts_import_root_aliases=base:. proto/service.proto
protoc -I=proto proto/service.proto \
    --js_out=import_style=commonjs:ts-pb \
    --grpc-web_out=import_style=commonjs,mode=grpcwebtext:ts-pb


```
## Compile code
```shell
python -m grpc_tools.protoc --proto_path=./proto --python_out=./client_python --grpc_python_out=./client_python proto/service.proto


python -m grpc_tools.protoc --proto_path=proto --python_out=./client_python --grpc_python_out=./client_python google/api/annotations.proto google/api/http.proto
```
## Using CLI
1. Start server
```shell
go run main.go --port 1222
```
2. Run client with username, default: Anonymous
```shell
go run client/main.go -N username -p 1222 -h localhost
```

## GRPC Call
1. Start server
```shell
go run main.go --port 1222
```
2. Dial to server
```go
conn, err := grpc.Dial(fmt.Sprintf("%v:%d", *host, *port), grpc.WithInsecure())
```
3. Create broadcastClient
```go
client = proto.NewBroadcastClient(conn)
```
4. Init a User object:
```go
	user := &proto.User{
		Id:          hex.EncodeToString(id[:]),
		DisplayName: *name,
	}
```
5. Connect User to stream and create event loop listen for new message :
```go
stream, err := client.CreateStream(context.Background(), &proto.Connect{
    User:   user,
    Active: true,
})

go func(str proto.Broadcast_CreateStreamClient) {
    defer wait.Done()
    
    for {
        msg, err := str.Recv()
        
        if err != nil {
            streamError = fmt.Errorf("Error reading message: %v", err)
            break
        }
        
        fmt.Printf("%v : %s\n", msg.User.DisplayName, msg.Message)
    }
}(stream)
```
6. When user submit chat call *BroadcastMessage* method:
```go
_, err := client.BroadcastMessage(context.Background(), msg)
```

protoc -I proto proto/service.proto \
--plugin="protoc-gen-ts=chat-ui/node_modules/bin/protoc-gen-ts" \
--ts_out="ts-pb"