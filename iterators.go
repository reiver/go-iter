package iter

import (
	"io"
)

// Iterators represents something that has a series of iterators.
//
// There are no more iterators when NextIterator() return nil (for Iterator).
//
// Iterators is used by FlattenedIterator.
type Iterators interface {
	io.Closer
	NextIterator()(Iterator,error)
}
