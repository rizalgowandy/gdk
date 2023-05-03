package natsx

import "context"

type (
	// SubscriberHandler is the handler definition to run a subscriber.
	SubscriberHandler func(ctx context.Context, subscriber *SubscriberImpl) error

	// SubscriberInterceptor is the middleware that will be executed before the current handler.
	SubscriberInterceptor func(ctx context.Context, subscriber *SubscriberImpl, handler SubscriberHandler) error
)

// SubscriberChain returns a single interceptor from multiple interceptors.
func SubscriberChain(interceptors ...SubscriberInterceptor) SubscriberInterceptor {
	n := len(interceptors)

	return func(ctx context.Context, subscriber *SubscriberImpl, handler SubscriberHandler) error {
		chainer := func(currentInter SubscriberInterceptor, currentHandler SubscriberHandler) SubscriberHandler {
			return func(currentCtx context.Context, currentSubscriber *SubscriberImpl) error {
				return currentInter(currentCtx, currentSubscriber, currentHandler)
			}
		}

		chainedHandler := handler
		for i := n - 1; i >= 0; i-- {
			chainedHandler = chainer(interceptors[i], chainedHandler)
		}

		return chainedHandler(ctx, subscriber)
	}
}
