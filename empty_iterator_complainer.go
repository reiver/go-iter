package iter

// EmptyIteratorComplainer is a type of error will be returned in situations where
// an empty iterator is not expected.
//
// Example:
//
//	err := iter.First{Iterator: iterator}.Decode(&datum)
//	if nil != err {
//		switch {
//		case iter.EmptyIteratorComplainer:
//			//@TODO
//		default:
//			return err
//		}
//	}
type EmptyIteratorComplainer interface {
	error
	EmptyIteratorComplainer()
}

type internalEmptyIteratorComplainer struct{}

func (receiver internalEmptyIteratorComplainer) Error() string {
	return "Empty Iterator"
}

func (receiver internalEmptyIteratorComplainer) EmptyIteratorComplainer() {

}
