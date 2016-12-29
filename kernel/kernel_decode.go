package iterkernel

import (
	"database/sql"
	"fmt"
	"time"
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
	}

	// According to the docs for sql.Scanner: https://golang.org/pkg/database/sql/#Scanner
	//
	//	"""
	//		The src value will be of one of the following types:
	//		
	//		int64
	//		float64
	//		bool
	//		[]byte
	//		string
	//		time.Time
	//		nil - for NULL values
	//	"""
	//
	// So, if the datum is one of those types, then we send it to Scan.
	//
	// If the datum is not one of those types, BUT can be safely converted
	// to one of those types (ex: uint8 -> int64), then we do the conversion
	// and then send that to Scan.
	switch p := x.(type) {
	case sql.Scanner:
		switch t := receiver.datum.(type) {
		case int64, float64, bool, []byte, string, time.Time:
			return p.Scan(t)
		case float32:
			return p.Scan( float64(t) )
		case int8:
			return p.Scan( int64(t) )
		case int16:
			return p.Scan( int64(t) )
		case int32:
			return p.Scan( int64(t) )
		case uint8:
			return p.Scan( int64(t) )
		case uint16:
			return p.Scan( int64(t) )
		case uint32:
			return p.Scan( int64(t) )
		}
	}


	return &internalBadTypeComplainer{
		actualType: fmt.Sprintf("%T", x),
		sourceType: fmt.Sprintf("%T", receiver.datum),
	}

}
