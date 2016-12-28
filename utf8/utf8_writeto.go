package iterutf8

import (
	"io"
	"unicode/utf8"
)

func (receiver *Slice) WriteTo(w io.Writer) (int64, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}

	var b [utf8.UTFMax]byte
	p := b[:]

	m := utf8.EncodeRune(p, receiver.datum)
	p = p[:m]

	n, err := w.Write(p)
	n64 := int64(n)

	return n64, err
}
