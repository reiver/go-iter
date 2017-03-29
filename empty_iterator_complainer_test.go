package iter

import (
	"testing"
)

func TestEmptyIteratorComplainer(t *testing.T) {

	var complainer EmptyIteratorComplainer = errEmptyIterator // THIS IS THE LINE THAT ACTUALLY MATTERS.

	if nil == complainer {
		t.Errorf("This should never happen.")
	}
}

func TestEmptyIteratorComplainerError(t *testing.T) {

	var err error = errEmptyIterator // THIS IS THE LINE THAT ACTUALLY MATTERS.

	if nil == err {
		t.Errorf("This should never happen.")
	}
}
