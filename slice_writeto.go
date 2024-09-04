package iter

import (
	"io"

	"github.com/reiver/go-erorr"
)

func (receiver *Slice[T]) WriteTo(w io.Writer) (int64, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}

	var n int
	var err error

	var value any = receiver.datum

	switch casted := value.(type) {
	case string:
		n, err = io.WriteString(w, casted)
	case []byte:
		n, err = w.Write(casted)
	case byte:
		var buffer [1]byte = [1]byte{casted}
		n, err = w.Write(buffer[:])
	default:
		return 0, erorr.Errorf("iter: cannot write something of type %T", receiver.datum)
	}

	n64 := int64(n)
	return n64, err
}
