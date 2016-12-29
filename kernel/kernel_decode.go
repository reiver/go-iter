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

	if nil != fn {
		decoded, err := fn(x)
		if nil != err {
			return err
		}
		if decoded {
			return nil
		}
	}

	switch t := receiver.datum.(type) {
	case float32:
		switch p := x.(type) {
		case *float32:
			if nil == p {
				return nil
			}
			*p = float32(t)
			return nil
		case *float64:
			if nil == p {
				return nil
			}
			*p = float64(t)
			return nil
		}
	case float64:
		switch p := x.(type) {
		case *float64:
			if nil == p {
				return nil
			}
			*p = float64(t)
			return nil
		}
	case int8:
		switch p := x.(type) {
		case *float32:
			if nil == p {
				return nil
			}
			*p = float32(t)
			return nil
		case *float64:
			if nil == p {
				return nil
			}
			*p = float64(t)
			return nil
		case *int8:
			if nil == p {
				return nil
			}
			*p = int8(t)
			return nil
		case *int16:
			if nil == p {
				return nil
			}
			*p = int16(t)
			return nil
		case *int32:
			if nil == p {
				return nil
			}
			*p = int32(t)
			return nil
		case *int64:
			if nil == p {
				return nil
			}
			*p = int64(t)
			return nil
		}
	case int16:
		switch p := x.(type) {
		case *float32:
			if nil == p {
				return nil
			}
			*p = float32(t)
			return nil
		case *float64:
			if nil == p {
				return nil
			}
			*p = float64(t)
			return nil
		case *int16:
			if nil == p {
				return nil
			}
			*p = int16(t)
			return nil
		case *int32:
			if nil == p {
				return nil
			}
			*p = int32(t)
			return nil
		case *int64:
			if nil == p {
				return nil
			}
			*p = int64(t)
			return nil
		}
	case int32:
		switch p := x.(type) {
		case *int32:
			if nil == p {
				return nil
			}
			*p = int32(t)
			return nil
		case *int64:
			if nil == p {
				return nil
			}
			*p = int64(t)
			return nil
		}
	case int64:
		switch p := x.(type) {
		case *int64:
			if nil == p {
				return nil
			}
			*p = int64(t)
			return nil
		}
	case uint8:
		switch p := x.(type) {
		case *float32:
			if nil == p {
				return nil
			}
			*p = float32(t)
			return nil
		case *float64:
			if nil == p {
				return nil
			}
			*p = float64(t)
			return nil
		case *int16:
			if nil == p {
				return nil
			}
			*p = int16(t)
			return nil
		case *int32:
			if nil == p {
				return nil
			}
			*p = int32(t)
			return nil
		case *int64:
			if nil == p {
				return nil
			}
			*p = int64(t)
			return nil
		case *uint8:
			if nil == p {
				return nil
			}
			*p = uint8(t)
			return nil
		case *uint16:
			if nil == p {
				return nil
			}
			*p = uint16(t)
			return nil
		case *uint32:
			if nil == p {
				return nil
			}
			*p = uint32(t)
			return nil
		case *uint64:
			if nil == p {
				return nil
			}
			*p = uint64(t)
			return nil
		}
	case uint16:
		switch p := x.(type) {
		case *float32:
			if nil == p {
				return nil
			}
			*p = float32(t)
			return nil
		case *float64:
			if nil == p {
				return nil
			}
			*p = float64(t)
			return nil
		case *int32:
			if nil == p {
				return nil
			}
			*p = int32(t)
			return nil
		case *int64:
			if nil == p {
				return nil
			}
			*p = int64(t)
			return nil
		case *uint16:
			if nil == p {
				return nil
			}
			*p = uint16(t)
			return nil
		case *uint32:
			if nil == p {
				return nil
			}
			*p = uint32(t)
			return nil
		case *uint64:
			if nil == p {
				return nil
			}
			*p = uint64(t)
			return nil
		}
	case uint32:
		switch p := x.(type) {
		case *float32:
			if nil == p {
				return nil
			}
			*p = float32(t)
			return nil
		case *float64:
			if nil == p {
				return nil
			}
			*p = float64(t)
			return nil
		case *int64:
			if nil == p {
				return nil
			}
			*p = int64(t)
			return nil
		case *uint32:
			if nil == p {
				return nil
			}
			*p = uint32(t)
			return nil
		case *uint64:
			if nil == p {
				return nil
			}
			*p = uint64(t)
			return nil
		}
	case uint64:
		switch p := x.(type) {
		case *uint64:
			if nil == p {
				return nil
			}
			*p = uint64(t)
			return nil
		}
	}

	switch p := x.(type) {
	case *interface{}:
		if nil == p {
			return nil
		}

		*p = receiver.datum

		return nil
	case sql.Scanner:
//@TODO: Should this be converted into the types that sql.Scanner supports?
		return p.Scan(receiver.datum)
	default:
		return &internalBadTypeComplainer{
			actualType: fmt.Sprintf("%T", p),
			sourceType: fmt.Sprintf("%T", receiver.datum),
		}
	}

	return nil
}
