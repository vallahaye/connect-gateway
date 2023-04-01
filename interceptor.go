package connectgateway

import "github.com/bufbuild/connect-go"

type chain struct {
	interceptors []connect.Interceptor
}

func newChain(interceptors []connect.Interceptor) *chain {
	var chain chain
	for i := len(interceptors) - 1; i >= 0; i-- {
		if interceptor := interceptors[i]; interceptor != nil {
			chain.interceptors = append(chain.interceptors, interceptor)
		}
	}
	return &chain
}

func (c *chain) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	for _, interceptor := range c.interceptors {
		next = interceptor.WrapUnary(next)
	}
	return next
}

func (c *chain) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	for _, interceptor := range c.interceptors {
		next = interceptor.WrapStreamingClient(next)
	}
	return next
}

func (c *chain) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	for _, interceptor := range c.interceptors {
		next = interceptor.WrapStreamingHandler(next)
	}
	return next
}
