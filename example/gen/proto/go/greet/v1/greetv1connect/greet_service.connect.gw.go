// Code generated by protoc-gen-connect-gateway. DO NOT EDIT.
//
// Source: greet/v1/greet_service.proto

package greetv1connect

import (
	context "context"
	fmt "fmt"
	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	connect_gateway "go.vallahaye.net/connect-gateway"
	v1 "go.vallahaye.net/connect-gateway/example/gen/proto/go/greet/v1"
)

// GreetServiceGatewayServer implements the gRPC server API for the GreetService service.
type GreetServiceGatewayServer struct {
	v1.UnimplementedGreetServiceServer
	greet connect_gateway.UnaryHandler[v1.GreetRequest, v1.GreetResponse]
}

// NewGreetServiceGatewayServer constructs a Connect-Gateway gRPC server for the GreetService
// service.
func NewGreetServiceGatewayServer(svc GreetServiceHandler, opts ...connect_gateway.HandlerOption) *GreetServiceGatewayServer {
	return &GreetServiceGatewayServer{
		greet: connect_gateway.NewUnaryHandler(GreetServiceGreetProcedure, svc.Greet, opts...),
	}
}

func (s *GreetServiceGatewayServer) Greet(ctx context.Context, req *v1.GreetRequest) (*v1.GreetResponse, error) {
	return s.greet(ctx, req)
}

// RegisterGreetServiceHandlerGatewayServer registers the Connect handlers for the GreetService
// "svc" to "mux".
func RegisterGreetServiceHandlerGatewayServer(mux *runtime.ServeMux, svc GreetServiceHandler, opts ...connect_gateway.HandlerOption) {
	if err := v1.RegisterGreetServiceHandlerServer(context.TODO(), mux, NewGreetServiceGatewayServer(svc, opts...)); err != nil {
		panic(fmt.Errorf("connect-gateway: %w", err))
	}
}
