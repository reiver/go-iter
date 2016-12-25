package iter

import (
	"fmt"
	"sync"
)

type Bools struct {
	Slice []bool
	err error
	index int
	datum bool
	mutex sync.RWMutex
}

func (receiver *Bools) Close() error {
	return nil
}

func (receiver *Bools) Decode(x interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	p, ok := x.(*bool)
	if !ok {
		return &internalBadTypeComplainer{fmt.Sprintf("%T", p)}
	}
	if nil == p {
		return nil
	}

	*p = receiver.datum

	return nil
}

func (receiver *Bools) Err() error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	return receiver.err
}

func (receiver *Bools) Next() bool {
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
