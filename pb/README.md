#pb

Generate `pb.go` file.

```
protoc -I . --go_out=plugins=grpc:. ./timestamp.proto
```