// Package local provides a local runtime
package local

import (
	"github.com/DragonFishYL/letui-micro/runtime"
)

// NewRuntime returns a new local runtime
func NewRuntime(opts ...runtime.Option) runtime.Runtime {
	return runtime.NewRuntime(opts...)
}
