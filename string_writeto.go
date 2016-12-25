package iter

import (
	"io"
)

func (receiver *String) WriteTo(w io.Writer) (int64, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}

	n, err := io.WriteString(w, receiver.datum)
	n64 := int64(n)

	return n64, err
}
