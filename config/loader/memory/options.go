package memory

import (
	"github.com/DragonFishYL/letui-micro/config/loader"
	"github.com/DragonFishYL/letui-micro/config/reader"
	"github.com/DragonFishYL/letui-micro/config/source"
)

// WithSource appends a source to list of sources
func WithSource(s source.Source) loader.Option {
	return func(o *loader.Options) {
		o.Source = append(o.Source, s)
	}
}

// WithReader sets the config reader
func WithReader(r reader.Reader) loader.Option {
	return func(o *loader.Options) {
		o.Reader = r
	}
}

func WithWatcherDisabled() loader.Option {
	return func(o *loader.Options) {
		o.WithWatcherDisabled = true
	}
}
