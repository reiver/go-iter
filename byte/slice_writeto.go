package iterbyte

import (
	"io"
)

func (receiver *Slice) WriteTo(w io.Writer) (int64, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}

	var b [1]byte
	b[0] = receiver.datum
	p := b[:]

	n, err := w.Write(p)
	n64 := int64(n)

	return n64, err
}
