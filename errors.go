package iter

import (
	"github.com/reiver/go-erorr"
)

var (
	errClosed            = erorr.Error("Closed")
	errDone              = erorr.Error("iter: done")
	errEmptyIterator     = internalEmptyIteratorComplainer{}
	errInternalError     = erorr.Error("Internal Error")
	errNilIterator       = erorr.Error("Nil Iterator")
	errNilIterators      = erorr.Error("iter: nil iterators")
	errNilReceiver       = erorr.Error("Nil Receiver")
	errNotFunc           = erorr.Error("Not Func")
	errTooManyIterations = erorr.Error("Too Many Iterations")
)
