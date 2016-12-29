package iterfloat64

import (
	"github.com/reiver/go-iter/kernel"

	"reflect"
)

type Slice struct {
	Slice []float64
	kernel iterkernel.Kernel
}

func (receiver *Slice) Close() error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.kernel.KernelClose()
}

// Err returns the error, if an error was encountered during an iteration.
// If no error was encountered during an iteration, then Err returns nil.
func (receiver *Slice) Err() error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.kernel.KernelErr()
}

func (receiver *Slice) Type() reflect.Type {
	return reflect.TypeOf((*float64)(nil)).Elem()
}
