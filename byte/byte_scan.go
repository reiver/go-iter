package iterbyte

import (
	"fmt"
)

func (receiver *Slice) Scan(dest ...interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	// Do NOT lock the mutex, since Decode will lock it.

	if expected, actual := 1, len(dest); expected != actual {
		return fmt.Errorf("iter: expected %d destination arguments in Scan, not %d", expected, actual)
	}

	dest0 := dest[0]

	if err := receiver.Decode(dest0); nil != err {
		return err
	}

	return nil
}
