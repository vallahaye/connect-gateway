module go.vallahaye.net/connect-gateway/example

go 1.18

require (
	connectrpc.com/connect v1.11.0
	github.com/go-chi/chi/v5 v5.0.8
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2
	go.vallahaye.net/connect-gateway v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.9.0
	google.golang.org/genproto/googleapis/api v0.0.0-20230525234035-dd9d682886f9
	google.golang.org/grpc v1.57.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.3.0
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/golang/glog v1.1.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/kr/text v0.2.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230526161137-0005af68ea54 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230525234030-28d5490b6b19 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace go.vallahaye.net/connect-gateway => ../
