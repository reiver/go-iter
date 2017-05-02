package iter

// First is useful when you only expect one (1) item in the iterator.
//
// Note that it will return an error if there is more than one (1) item
// in the iterator. (It checks for this.)
//
// Example Usage
//
//	iterator := ...
//	
//	err := iter.First{Iterator: iterator}.Decode(v)
//	if nil != err {
//		switch {
//		case iter.EmptyIteratorComplainer:
//			//@TODO
//		default:
//			return err
//		}
//	}
type First struct {
	Iterator Iterator
}

func (receiver First) Decode(v interface{}) error {
	return receiver.decode(v)
}

func (receiver First) decode(v interface{}) (err error) {

	iterator := receiver.Iterator
	if nil == iterator {
		return errNilIterator
	}
	defer func(i Iterator){
		if e := i.Close(); nil != e {
			if nil == err {
				err = e
			}
		}
	}(iterator)

	if ! iterator.Next() {
		return errEmptyIterator
	}
	if err := iterator.Decode(v); nil != err {
		return err
	}
	if err := iterator.Err(); nil != err {
		return err
	}

	return nil
}
