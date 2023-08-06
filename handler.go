package connectgateway

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type handlerConfig struct {
	Interceptor connect.Interceptor
}

type UnaryHandler[Req, Res any] func(context.Context, *Req) (*Res, error)

func NewUnaryHandler[Req, Res any](
	procedure string,
	unary func(context.Context, *connect.Request[Req]) (*connect.Response[Res], error),
	opts ...HandlerOption,
) UnaryHandler[Req, Res] {
	untypedUnary := func(ctx context.Context, untypedReq connect.AnyRequest) (connect.AnyResponse, error) {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		req, ok := untypedReq.(*connect.Request[Req])
		if !ok {
			return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("unexpected handler request type %T", req))
		}
		res, err := unary(ctx, req)
		if res == nil && err == nil {
			// This is going to panic during serialization. Debugging is much easier
			// if we panic here instead, so we can include the procedure name.
			panic(fmt.Sprintf("%s returned nil *connect.Response and nil error", procedure)) //nolint: forbidigo
		}
		return res, err
	}
	var config handlerConfig
	for _, opt := range opts {
		opt.applyToHandler(&config)
	}
	if interceptor := config.Interceptor; interceptor != nil {
		untypedUnary = interceptor.WrapUnary(untypedUnary)
	}
	return func(ctx context.Context, msg *Req) (*Res, error) {
		req := connect.NewRequest(msg)
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			for k := range md {
				for _, v := range md.Get(k) {
					req.Header().Add(k, v)
				}
			}
		}
		untypedRes, err := untypedUnary(ctx, req)
		if err != nil {
			return nil, status.Error(codes.Code(connect.CodeOf(err)), err.Error())
		}
		res, ok := untypedRes.(*connect.Response[Res])
		if !ok {
			return nil, status.Errorf(codes.Internal, "unexpected handler response type %T", res)
		}
		if h := res.Header(); len(h) > 0 {
			if err := grpc.SendHeader(ctx, metadataFromHeader(h)); err != nil {
				return nil, err
			}
		}
		if h := res.Trailer(); len(h) > 0 {
			if err := grpc.SetTrailer(ctx, metadataFromHeader(h)); err != nil {
				return nil, err
			}
		}
		return res.Msg, nil
	}
}
