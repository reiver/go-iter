package iterfloat64

// Decode stores the next datum in the data represented by the empty interface value `x`.
// If `x` is nil, the value will be discarded.
// Otherwise, the value underlying `x` must be a pointer to the correct type for the next datum.
func (receiver *Slice) Decode(x interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.nucleus._decode(receiver.decode, x)
}

func (receiver *Slice) decode(x interface{}) (bool, error) {
	if nil == receiver {
		return false, errNilReceiver
	}

	if nil == x {
		return false, nil
	}

	f64, err := receiver.nucleus._datum()
	if nil != err {
		return false, err
	}

	switch p := x.(type) {
	case *float64:
                if nil == p {
                        return true, nil
                }

                *p = f64

		return true, nil
	default:
                return false, nil
	}
}
