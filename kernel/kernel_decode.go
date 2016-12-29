package iterkernel

import (
	"database/sql"
	"fmt"
)

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
//@TODO: Should this be converted into the types that sql.Scanner supports?
		return p.Scan(receiver.datum)
	default:
		return &internalBadTypeComplainer{fmt.Sprintf("%T", p)}
	}

	return nil
}
