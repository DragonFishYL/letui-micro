package router

import (
	"github.com/DragonFishYL/letui-micro/api/resolver"
	"github.com/DragonFishYL/letui-micro/api/resolver/vpath"
	"github.com/DragonFishYL/letui-micro/logger"
	"github.com/DragonFishYL/letui-micro/registry"
)

type Options struct {
	Handler  string
	Registry registry.Registry
	Resolver resolver.Resolver
	Logger   logger.Logger
}

type Option func(o *Options)

func NewOptions(opts ...Option) Options {
	options := Options{
		Handler:  "meta",
		Registry: registry.DefaultRegistry,
		Logger:   logger.DefaultLogger,
	}

	for _, o := range opts {
		o(&options)
	}

	if options.Resolver == nil {
		options.Resolver = vpath.NewResolver(
			resolver.WithHandler(options.Handler),
		)
	}

	return options
}

func WithHandler(h string) Option {
	return func(o *Options) {
		o.Handler = h
	}
}

func WithRegistry(r registry.Registry) Option {
	return func(o *Options) {
		o.Registry = r
	}
}

func WithResolver(r resolver.Resolver) Option {
	return func(o *Options) {
		o.Resolver = r
	}
}

// WithLogger sets the underline logger
func WithLogger(l logger.Logger) Option {
	return func(o *Options) {
		o.Logger = l
	}
}
