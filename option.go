package connectgateway

import "github.com/bufbuild/connect-go"

type HandlerOption interface {
	applyToHandler(*handlerConfig)
}

type interceptorsOption struct {
	Interceptors []connect.Interceptor
}

func (o *interceptorsOption) applyToHandler(config *handlerConfig) {
	config.Interceptor = o.chainWith(config.Interceptor)
}

func (o *interceptorsOption) chainWith(interceptor connect.Interceptor) connect.Interceptor {
	if len(o.Interceptors) == 0 {
		return interceptor
	}
	if interceptor == nil && len(o.Interceptors) == 1 {
		return o.Interceptors[0]
	}
	if interceptor == nil && len(o.Interceptors) > 1 {
		return newChain(o.Interceptors)
	}
	return newChain(append([]connect.Interceptor{interceptor}, o.Interceptors...))
}

func WithInterceptors(interceptors ...connect.Interceptor) HandlerOption {
	return &interceptorsOption{interceptors}
}
