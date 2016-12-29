package iterkernel

import (
	"database/sql"
	"fmt"
	"sync"
)

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

func (receiver *Kernel) KernelDecode(fn func(interface{})(bool,error), x interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	{
		decoded, err := fn(x)
		if nil != err {
			return err
		}
		if decoded {
			return nil
		}
	}

	switch p := x.(type) {
	case *interface{}:
		if nil == p {
			return nil
		}

		*p = receiver.datum
	case sql.Scanner:
		return p.Scan(receiver.datum)
	default:
		return &internalBadTypeComplainer{fmt.Sprintf("%T", p)}
	}

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
