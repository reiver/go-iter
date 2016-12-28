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

func (receiver *Kernel) KernelErr() error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	return receiver.err
}

func (receiver *Kernel) KernelNext(fn func(int,interface{})(bool, error)) bool {
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

	index := receiver.index

	var datum interface{}

	next, err := fn(index, &datum)
	if nil != err {
		receiver.err = err
		return false
	}
	if !next {
		return false
	}

	receiver.datum = datum
	receiver.index++

	return true
}
