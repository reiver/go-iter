package iter

import (
	"errors"

	"github.com/reiver/go-erorr"
)

// FlattenedIterator flattens Iterators into a "normal" Iterator.
type FlattenedIterator struct {
	Iterators Iterators

	err error
	iterator Iterator
	closed bool
	done bool
}

var _ Iterator = &FlattenedIterator{}

func (receiver *FlattenedIterator) Close() error {
	if nil == receiver {
		return errNilReceiver
	}

	if receiver.closed {
		return nil
	}

	var err1 error
	{
		var iterator Iterator = receiver.iterator

		if nil != iterator {
			err1 = iterator.Close()
		}
	}

	var err2 error
	{
		var iterators Iterators = receiver.Iterators

		if nil != iterators {
			err2 = iterators.Close()
		}
	}

	receiver.closed = true
	var err error = errors.Join(err1, err2)
	return err
}

func (receiver *FlattenedIterator) Decode(dst any) error {
	if nil == receiver {
		return errNilReceiver
	}

	if receiver.done {
		return errDone
	}
	if receiver.closed {
		return errClosed
	}

	if nil != receiver.err {
		return erorr.Errorf("iter: cannot decode because of previous error: %w", receiver.err)
	}

	var iterator Iterator = receiver.iterator
	if nil == iterator {
		return errNilIterator
	}

	return iterator.Decode(dst)
}

func (receiver *FlattenedIterator) Err() error {
	if nil == receiver {
		return errNilReceiver
	}

	var err1 error
	{
		var iterator Iterator = receiver.iterator
		if nil != iterator {
			err1 = iterator.Err()
		}
	}

	var err2 error
	{
		err2 = receiver.err
	}

	return errors.Join(err1, err2)
}

func (receiver *FlattenedIterator) Next() bool {
	if nil == receiver {
		return false
	}

	if receiver.done {
		return false
	}
	if receiver.closed {
		return false
	}

	var iterator Iterator = receiver.iterator
	{
		if nil == iterator {
			if !receiver.nextIterator() {
				return false
			}
			iterator = receiver.iterator
			if nil == iterator {
				return false
			}
		}
	}

	var result bool = iterator.Next()
	if false == result {
		if !receiver.nextIterator() {
			return false
		}
		iterator = receiver.iterator
		if nil == iterator {
			return false
		}
		result = iterator.Next()
	}


	return result
}

func (receiver *FlattenedIterator) nextIterator() bool {
	if nil == receiver {
		return false
	}

	{
		var iterators Iterators = receiver.Iterators
		if nil == iterators {
			receiver.err = errNilIterators
			return false
		}

		var iterator Iterator
		var err error
		iterator, err = iterators.NextIterator()
		if nil != err {
			receiver.err = err
			return false
		}
		if nil == iterator {
			receiver.done = true
			return false
		}
		receiver.iterator = iterator
	}

	return true
}
