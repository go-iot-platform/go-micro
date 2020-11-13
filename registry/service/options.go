package service

import (
	"context"

	"github.com/panovateam/go-micro/client"
	"github.com/panovateam/go-micro/registry"
)

type clientKey struct{}

// WithClient sets the RPC client
func WithClient(c client.Client) registry.Option {
	return func(o *registry.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}

		o.Context = context.WithValue(o.Context, clientKey{}, c)
	}
}
