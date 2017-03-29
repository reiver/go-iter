package iter

import (
	"errors"
)

var (
	errClosed            = errors.New("Closed")
	errEmptyIterator     = internalEmptyIteratorComplainer{}
	errInternalError     = errors.New("Internal Error")
	errNilIterator       = errors.New("Nil Iterator")
	errNilReceiver       = errors.New("Nil Receiver")
	errNotFunc           = errors.New("Not Func")
	errTooManyIterations = errors.New("Too Many Iterations")
)
