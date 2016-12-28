package iterfloat64x2

import (
	"fmt"
)

func (receiver *Slice) Scan(dest ...interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	// Do NOT lock the mutex, since Decode will lock it.

	if expected, actual := 2, len(dest); expected != actual {
		return fmt.Errorf("iter: expected %d destination arguments in Scan, not %d", expected, actual)
	}

	if err := receiver.Decode(dest); nil != err {
		return err
	}

	return nil
}
