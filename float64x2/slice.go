package iterfloat64x2

import (
	"database/sql"
	"fmt"
	"sync"
)

type Slice struct {
	Slice [][2]float64
	err error
	index int
	mutex sync.RWMutex
	closed bool
	datum [2]float64
}

func (receiver *Slice) Close() error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	receiver.closed = true

	return nil
}

// Decode stores the next datum in the data represented by the empty interface value `x`.
// If `x` is nil, the value will be discarded.
// Otherwise, the value underlying `x` must be a pointer to the correct type for the next datum.
func (receiver *Slice) Decode(x interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	switch p := x.(type) {
	case *[2]float64:
		if nil == p {
			return nil
		}

		*p = receiver.datum
	case *interface{}:
		if nil == p {
			return nil
		}

		*p = receiver.datum
	case []interface{}:
		if nil == p {
			return nil
		}

		if expected, actual := 2, len(p); expected != actual {
			return fmt.Errorf("iter: expected %d destination arguments in Scan, not %d", expected, actual)
		}

		switch f := p[0].(type) {
		case *float64:
			*f = receiver.datum[0]
		case *interface{}:
			*f = receiver.datum[0]
		case sql.Scanner:
			if err := f.Scan(receiver.datum[0]); nil != err {
				return err
			}
		default:
			return &internalBadTypeComplainer{fmt.Sprintf("%T", f)}
		}

		switch f := p[1].(type) {
		case *float64:
			*f = receiver.datum[1]
		case *interface{}:
			*f = receiver.datum[1]
		case sql.Scanner:
			if err := f.Scan(receiver.datum[1]); nil != err {
				return err
			}
		default:
			return &internalBadTypeComplainer{fmt.Sprintf("%T", f)}
		}

	default:
		return &internalBadTypeComplainer{fmt.Sprintf("%T", p)}
	}

	return nil
}

// Err returns the error, if an error was encountered during an iteration.
// If no error was encountered during an iteration, then Err returns nil.
func (receiver *Slice) Err() error {
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
func (receiver *Slice) Next() bool {
	if nil == receiver {
		return false
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if nil != receiver.err {
		return false
	}

	if receiver.closed {
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
