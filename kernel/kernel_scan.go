package iterkernel

import (
	"fmt"
)

func (receiver *Kernel) KernelScan(fn func(interface{})(bool,error), dest ...interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	// Do NOT lock the mutex, since _decode will lock it.

	if expected, actual := 1, len(dest); expected != actual {
		return fmt.Errorf("iter: expected %d destination arguments in Scan, not %d", expected, actual)
	}

	dest0 := dest[0]

	if err := receiver.KernelDecode(fn, dest0); nil != err {
		return err
	}

	return nil
}
