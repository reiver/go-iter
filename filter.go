package iter

import (
	"github.com/reiver/go-erorr"
)

type Filter[T any] struct {
	Iterator Iterator
	Func func(T)bool

	decodedValue T
	decodedError error
}

var _ Iterator = &Filter[string]{}

func (receiver *Filter[T]) Close() error {
	var iterator Iterator = receiver.Iterator
	if nil == iterator {
		return nil
	}

	return iterator.Close()
}

func (receiver *Filter[T]) Decode(dst any) error {
	var err error = receiver.decodedError
	if nil != err {
		return err
	}

	ptr, casted := dst.(*T)
	if !casted {
		return erorr.Errorf("iter: cannot decode something of type %T into type %T, must be of type *%T", receiver.decodedValue, dst, receiver.decodedValue)
	}

	*ptr = receiver.decodedValue
	return nil
}

func (receiver *Filter[T]) Err() error {
	var iterator Iterator = receiver.Iterator
	if nil == iterator {
		return nil
	}

	return iterator.Err()
}

func (receiver *Filter[T]) Next() bool {
	var iterator Iterator = receiver.Iterator
	if nil == iterator {
		return false
	}

	for iterator.Next() {

		err := iterator.Decode(&receiver.decodedValue)
		if nil != err{
			receiver.decodedError = err
			return true
		}

		var fn func(T)bool = receiver.Func
		if nil == fn {
			fn = truefunc
		}

		if fn(receiver.decodedValue) {
			return true
		}
	}

	return false
}
