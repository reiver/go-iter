package iter

type SplitIterators[T any] struct {
	Iterator Iterator
	Func func(T) (Iterator, error)
	closed bool
}

var _ Iterators = &SplitIterators[string]{}

func (receiver *SplitIterators[T]) Close() error {
	if nil == receiver {
		return errNilReceiver
	}

	if receiver.closed {
		return nil
	}

	var iterator Iterator = receiver.Iterator
	if nil == iterator {
		return errNilIterator
	}

	if err := iterator.Close(); nil != err {
		return err
	}
	receiver.closed = true
	receiver.Iterator = nil
	return nil
}

func (receiver *SplitIterators[T]) NextIterator() (Iterator, error) {
	if nil == receiver {
		return nil, errNilReceiver
	}

	if receiver.closed {
		return nil, nil
	}

	var iterator Iterator = receiver.Iterator
	if nil == iterator {
		return nil, errNilIteratorInsideSplitIterators
	}

	var function func(T)(Iterator,error) = receiver.Func
	if nil == function {
		return nil, errNilFunc
	}

	var value T
	{
		var next bool = iterator.Next()
		if !next {
			if err := iterator.Err(); nil != err {
				return nil, err
			}
			return nil, nil
		}

		if err := iterator.Decode(&value); nil != err {
			return nil, err
		}
	}

	return function(value)
}
