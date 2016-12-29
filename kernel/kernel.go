package iterkernel

import (
	"sync"
)

// Kernel deals with the "book keeping" needed when creating a common iterator.
//
// For example:
// making it so calling Close stops Next from being able to load another value;
// dealing with having an error that occurred when Next is called get returned  when Err is called;
// dealing with decoding & scanning into a *interface{},
// dealing with decoding & scanning into a sql.Scanner,
// dealing with the read-write mutex so that this can be safely accessed from multiple threads,
// etc.
type Kernel struct {
	err error
	index int
	mutex sync.RWMutex
	closed bool
	datum interface{}
}

func (receiver *Kernel) KernelClose() error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	receiver.closed = true

	return nil
}

// KernelErr returns any errors that occurred when KerneNext was called, if there are any.
// (There might not have been any.)
//
// It is the partner of the Next method found in common iterators:
//
//	Err() error
//
// The error could come from this kernel itself.
//
// Also, the error could come from the func passed to KernelNext.
func (receiver *Kernel) KernelErr() error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	return receiver.err
}
