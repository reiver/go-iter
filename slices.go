package iter

type Slices[T any] struct {
	Slices [][]T
}

var _ Iterators = &Slices[string]{}

func (receiver *Slices[T]) Close() error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.Slices = nil
	return nil
}

func (receiver *Slices[T]) nextSlice() ([]T, error) {
	if nil == receiver {
		return nil, errNilReceiver
	}

	var slices [][]T = receiver.Slices

	if len(slices) <= 0 {
		receiver.Slices = nil
		return nil, nil
	}

	var slice []T = slices[0]
	receiver.Slices = slices[1:]
	if len(receiver.Slices) <= 0 {
		receiver.Slices = nil
	}

	return slice, nil
}

func (receiver *Slices[T]) NextIterator() (Iterator, error) {
	if nil == receiver {
		return nil, errNilReceiver
	}

	slice, err := receiver.nextSlice()
	if nil != err {
		return nil, err
	}
	if len(slice) <= 0 {
		return nil, nil
	}

	var iterator Iterator = &Slice[T]{Slice:slice}

	return iterator, nil
}
