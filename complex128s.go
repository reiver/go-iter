package iter

import (
	"fmt"
	"sync"
)

type Complex128s struct {
	Slice []complex128
	err error
	index int
	datum complex128
	mutex sync.RWMutex
}

func (receiver *Complex128s) Close() error {
	return nil
}

// Decode stores the next datum in the data represented by the empty interface value `x`.
// If `x` is nil, the value will be discarded.
// Otherwise, the value underlying `x` must be a pointer to the correct type for the next datum.
func (receiver *Complex128s) Decode(x interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	switch p := x.(type) {
	case *complex128:
		if nil == p {
			return nil
		}

		*p = receiver.datum
	case *interface{}:
		if nil == p {
			return nil
		}

		*p = receiver.datum
	default:
		return &internalBadTypeComplainer{fmt.Sprintf("%T", p)}
	}

	return nil
}

// Err returns the error, if an error was encountered during an iteration.
// If no error was encountered during an iteration, then Err returns nil.
func (receiver *Complex128s) Err() error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	return receiver.err
}

// Next prepares the next datum for reading via the Decode method.
//
// If there is indeed a next datum, Next returns true.
//
// And a call to the Decode method will obtain that next datum.
//
// If there is no next datum, Next returns false.
//
// If an error was encountered while calling Next, Next will also return false.
//
// And a call to the Err method will obtain that error.
//
// Err should be called to distinguish between the two cases where the Next
// method can return false: 'no next datum' and 'an error was encountered'.
func (receiver *Complex128s) Next() bool {
	if nil == receiver {
		return false
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if nil != receiver.err {
		return false
	}

	slice := receiver.Slice
	if nil == slice  {
		return false
	}

	index := receiver.index
	if len(slice) <= index {
		return false
	}

	receiver.datum = slice[index]
	receiver.index++

	return true
}