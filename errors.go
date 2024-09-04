package iter

import (
	"github.com/reiver/go-erorr"
)

var (
	errClosed            = erorr.Error("iter: closed")
	errDone              = erorr.Error("iter: done")
	errEmptyIterator     = internalEmptyIteratorComplainer{}
	errInternalError     = erorr.Error("iter: internal error")
	errNilFunc           = erorr.Error("iter: nil func")
	errNilIteratorInsideSplitIterators = erorr.Error("iter: nil iterator inside iter.SplitIterators")
	errNilIterator       = erorr.Error("iter: nil iterator")
	errNilIterators      = erorr.Error("iter: nil iterators")
	errNilReceiver       = erorr.Error("iter: nil receiver")
	errNotFunc           = erorr.Error("iter: not func")
	errTooManyIterations = erorr.Error("iter: too many iterations")
)
