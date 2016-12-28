package iterfloat64

import (
	"github.com/reiver/go-iter/kernel"

	"math"
	"reflect"
)

type nucleus struct {
	kernel iterkernel.Kernel
}

func (receiver *nucleus) _close() error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.kernel.KernelClose()
}

func (receiver *nucleus) _datum() (float64, error) {
	if nil == receiver {
		return math.NaN(), errNilReceiver
	}

	datum, err := receiver.kernel.KernelDatum()
	if nil != err {
		return math.NaN(), err
	}

	f64, ok := datum.(float64)
	if !ok {
		return math.NaN(), errInternalError
	}

	return f64, nil
}

func (receiver *nucleus) _err() error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.kernel.KernelErr()
}

func (receiver *nucleus) _decode(fn func(interface{})(bool,error), x interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.kernel.KernelDecode(fn, x)
}

func (receiver *nucleus) _next(fn func(int,interface{})(bool, error)) bool {
	if nil == receiver {
		return false
	}

	return receiver.kernel.KernelNext(fn)

}

func (receiver *nucleus) _scan(fn func(interface{})(bool,error), dest ...interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.kernel.KernelScan(fn, dest...)
}

func (receiver *nucleus) _type() reflect.Type {
	return reflect.TypeOf((*float64)(nil)).Elem()
}
