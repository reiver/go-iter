package iter

import (
	"database/sql"
	"fmt"
	"sync"
	"unicode/utf8"
)

type UTF8 struct {
	Slice []byte
	err error
	index int
	mutex sync.RWMutex
	closed bool
	datum rune
}

func (receiver *UTF8) Close() error {
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
func (receiver *UTF8) Decode(x interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	switch p := x.(type) {
	case *rune:
		if nil == p {
			return nil
		}

		*p = receiver.datum
	case *interface{}:
		if nil == p {
			return nil
		}

		*p = receiver.datum
	case sql.Scanner:
		var buffer [utf8.UTFMax]byte
		ptr := buffer[:]

		n := utf8.EncodeRune(ptr, receiver.datum)
		ptr = ptr[:n]

		return p.Scan(ptr)
	default:
		return &internalBadTypeComplainer{fmt.Sprintf("%T", p)}
	}

	return nil
}

// Err returns the error, if an error was encountered during an iteration.
// If no error was encountered during an iteration, then Err returns nil.
func (receiver *UTF8) Err() error {
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
func (receiver *UTF8) Next() bool {
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

	r, n := utf8.DecodeRune(slice[index:])
	receiver.datum = r
	receiver.index += n

	return true
}