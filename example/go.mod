module go.vallahaye.net/connect-gateway/example

go 1.18

require (
	github.com/bufbuild/connect-go v1.5.2
	github.com/go-chi/chi/v5 v5.0.8
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2
	go.vallahaye.net/connect-gateway v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.7.0
	google.golang.org/genproto v0.0.0-20230306155012-7f2fa6fef1f4
	google.golang.org/grpc v1.53.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.3.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/golang/glog v1.0.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/kr/text v0.2.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace go.vallahaye.net/connect-gateway => ../
