package iterfloat64

import (
	"database/sql"
	"fmt"
	"sync"
)

type common struct {
	err error
	index int
	mutex sync.RWMutex
	closed bool
	datum interface{}
}

func (receiver *common) _close() error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	receiver.closed = true

	return nil
}

func (receiver *common) _decode(fn func(interface{})bool, x interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	if fn(x) {
		return nil
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

func (receiver *common) _err() error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	return receiver.err
}

func (receiver *common) _next(fn func(int,interface{})(bool, error)) bool {
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
