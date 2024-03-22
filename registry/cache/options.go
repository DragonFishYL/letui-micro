package cache

import (
	"time"

	"github.com/DragonFishYL/letui-micro/logger"
)

// WithTTL sets the cache TTL
func WithTTL(t time.Duration) Option {
	return func(o *Options) {
		o.TTL = t
	}
}

// WithLogger sets the underline logger
func WithLogger(l logger.Logger) Option {
	return func(o *Options) {
		o.Logger = l
	}
}
