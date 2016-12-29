package iterfloat64

type Slice struct {
	Slice []float64
	nucleus
}

func (receiver *Slice) Close() error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.nucleus._close()
}

// Err returns the error, if an error was encountered during an iteration.
// If no error was encountered during an iteration, then Err returns nil.
func (receiver *Slice) Err() error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.nucleus._err()
}
