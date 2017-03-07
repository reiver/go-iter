package iter

var (
	// Empty is an empty iterator.
	//
	// It is useful in cases where you want to quickly and cheaply return an empty iterator.
	//
	//
	// Usage:
	//
	//	func fn() (iter.Iterator, error) {
	//	
	//		// ...
	//	
	//		return iter.Empty, nil
	Empty Iterator
)

func init() {
	Empty = internalEmptyIterator{}
}

type internalEmptyIterator struct{}

func (receiver internalEmptyIterator) Close() error {
	return nil
}

func (receiver internalEmptyIterator) Decode(interface{}) error {
	return errClosed
}

func (receiver internalEmptyIterator) Err() error {
	return nil
}

func (receiver internalEmptyIterator) Next() bool {
	return false
}

