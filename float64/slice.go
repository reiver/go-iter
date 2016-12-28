package iterfloat64

import (
	"fmt"
)

type Slice struct {
	Slice []float64
	common
}

func (receiver *Slice) Close() error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.common._close()
}

// Decode stores the next datum in the data represented by the empty interface value `x`.
// If `x` is nil, the value will be discarded.
// Otherwise, the value underlying `x` must be a pointer to the correct type for the next datum.
func (receiver *Slice) Decode(x interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.common._decode(receiver.decode, x)
}

func (receiver *Slice) decode(x interface{}) bool {
	if nil == receiver {
		return false
	}

	if nil == x {
		return false
	}

	switch p := x.(type) {
	case *float64:
                if nil == p {
                        return true
                }

                *p = receiver.datum.(float64)

		return true
	default:
                return false
	}
}

// Err returns the error, if an error was encountered during an iteration.
// If no error was encountered during an iteration, then Err returns nil.
func (receiver *Slice) Err() error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.common._err()
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
func (receiver *Slice) Next() bool {
	if nil == receiver {
		return false
	}

	return receiver.common._next(receiver.next)
}

func (receiver *Slice) next(index int, v interface{}) (bool, error) {
	if nil == receiver {
		return false, errNilReceiver
	}

	slice := receiver.Slice
	if nil == slice  {
		return false, nil
	}

	if len(slice) <= index {
		return false, nil
	}

	datum := slice[index]

	switch t := v.(type) {
	case *interface{}:
		*t = datum
	default:
		return false, &internalBadTypeComplainer{fmt.Sprintf("%T", t)}
	}

	return true, nil
}
