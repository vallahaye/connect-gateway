package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	greetv1 "go.vallahaye.net/connect-gateway/example/gen/proto/go/greet/v1"
	"go.vallahaye.net/connect-gateway/example/gen/proto/go/greet/v1/greetv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	exampleSrv := newExampleServer(&greetServiceHandlerImpl{})
	if err := exampleSrv.Serve(lis); err != http.ErrServerClosed {
		log.Fatalf("failed to serve: %v", err)
	}
}

func newExampleServer(greetSvc greetv1connect.GreetServiceHandler) *http.Server {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(5 * time.Second))
	router.Mount(greetv1connect.NewGreetServiceHandler(greetSvc))
	router.Mount(newExampleGatewayHandler(greetSvc))
	return &http.Server{
		Handler:        h2c.NewHandler(router, &http2.Server{}),
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 10 * 1024,
	}
}

func newExampleGatewayHandler(greetSvc greetv1connect.GreetServiceHandler) (string, http.Handler) {
	mux := runtime.NewServeMux()
	greetv1connect.RegisterGreetServiceHandlerGatewayServer(mux, greetSvc)
	return "/", mux
}

type greetServiceHandlerImpl struct{}

func (s *greetServiceHandlerImpl) Greet(ctx context.Context, req *connect.Request[greetv1.GreetRequest]) (*connect.Response[greetv1.GreetResponse], error) {
	msg := &greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello %s!", req.Msg.Name),
	}
	return connect.NewResponse(msg), nil
}
