package internaltesting

import (
	"github.com/reiver/go-erorr"
)

type TestScanner[T any] struct {
	value T
	scanned bool
}

func (receiver TestScanner[T]) Scanned() bool {
	return receiver.scanned
}

func (receiver TestScanner[T]) Value() T {
	return receiver.value
}

func (receiver *TestScanner[T]) Scan(src any) error {
	if nil == src {
		return erorr.Errorf("internaltesting: cannot scan into nil: (%T) %v", src, src)
	}

	b, ok := src.(T)
	if !ok {
		return erorr.Errorf("internaltesting: cannot not cast %T into %T.", src, b)
	}

	receiver.value   = b
	receiver.scanned = true

	return nil
}
