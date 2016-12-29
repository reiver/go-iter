package iterkernel

// KernelNext is responsible for loading the next value for the iterator.
//
// It is the partner of the Next method found in common iterators:
//
//	Next() bool
//
// (KernelNext does all the "book keeping" for a Next method.)
//
// Like Next, KernelNext will return true if there is a next value,
// and will return false is there is not a next value.
//
// Also, like Next, KernelNext returning false could be because of two different reasons.
//
// #1 because there are no more values in the iterator.
//
// Or #2 because an error occurred while trying to obain the next value.
//
// (KernelErr needs to be called to determine which was the case.)
//
// The parameter passed to KernelNext is a func with the signature:
//
//	func(int, interface{}) (bool, error)
//
// This func is what actually (potentially) gives KernelNext the next value.
//
// KernelNext calls this func, and passes it #1 the iteration number
// (which starts counting a zero), and #2 a pointer to where the func
// should store the next value.
//
// If the func loaded a next value successfully, then it should return like the following:
//
//	return true, nil
//
// If an error occurred while the func was trying to load a next value, then it should return like the following:
//
//	return false, err
//
// If there is no next value (and no errors occurred) then it should return like the following:
//
//	return false, nil
//
// An example func might look like:
//
//	func (iterationNumber int, v interface{}) (bool, error) {
//		if 10 < index {
//			return false, nil
//		}
//		
//		x := float64(index)
//		
//		datum := x*x
//
//		p, ok := v.(*interface{})
//		if !ok {
//			return false, errInternalError
//		}
//		
//		*p = datum
//		
//		return true, nil
//	}
//
// That example would produce the sequence: 0.0, 1.0, 4.0, 9,0, 25,0, 36.0, 49.0, 64.0, 81.0.
func (receiver *Kernel) KernelNext(fn func(int,interface{})(bool, error)) bool {
	if nil == receiver {
		return false
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if nil != receiver.err {
		return false
	}

	if receiver.closed {
		return false
	}

	index := receiver.index

	var datum interface{}

	next, err := fn(index, &datum)
	if nil != err {
		receiver.err = err
		return false
	}
	if !next {
		return false
	}

	receiver.datum = datum
	receiver.index++

	return true
}
