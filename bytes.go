package iter

import (
	"fmt"
	"sync"
)

type Bytes struct {
	Slice []byte
	err error
	index int
	datum byte
	mutex sync.RWMutex
}

func (receiver *Bytes) Close() error {
	return nil
}

func (receiver *Bytes) Decode(x interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	p, ok := x.(*byte)
	if !ok {
		return &internalBadTypeComplainer{fmt.Sprintf("%T", p)}
	}
	if nil == p {
		return nil
	}

	*p = receiver.datum

	return nil
}

// Err returns the error, if an error was encountered during an iteration.
// If no error was encountered during an iteration, then Err returns nil.
func (receiver *Bytes) Err() error {
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
func (receiver *Bytes) Next() bool {
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
