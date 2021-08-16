package nsqx

import "context"

type (
	// ConsumerHandler is the handler definition to run a consumer.
	ConsumerHandler func(ctx context.Context, consumer *Consumer) error

	// ConsumerInterceptor is the middleware that will be executed before the current handler.
	ConsumerInterceptor func(ctx context.Context, consumer *Consumer, handler ConsumerHandler) error
)

// ConsumerChain returns a single interceptor from multiple interceptors.
func ConsumerChain(interceptors ...ConsumerInterceptor) ConsumerInterceptor {
	n := len(interceptors)

	return func(ctx context.Context, consumer *Consumer, handler ConsumerHandler) error {
		chainer := func(currentInter ConsumerInterceptor, currentHandler ConsumerHandler) ConsumerHandler {
			return func(currentCtx context.Context, currentConsumer *Consumer) error {
				return currentInter(currentCtx, currentConsumer, currentHandler)
			}
		}

		chainedHandler := handler
		for i := n - 1; i >= 0; i-- {
			chainedHandler = chainer(interceptors[i], chainedHandler)
		}

		return chainedHandler(ctx, consumer)
	}
}
