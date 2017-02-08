package iter

import (
	"github.com/reiver/go-iter/string"

	"testing"
)

func TestFirst(t *testing.T) {

	const expectedValue = "apple"

	iterator := &iterstring.Slice{
		Slice: []string{
			expectedValue,
		},
	}

	var datum string

	if err := (First{iterator}).Decode(&datum); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}
	if expected, actual := expectedValue, datum; expected != actual {
		t.Errorf("Expected %q, but actually got %q.", expected, actual)
		return
	}
}

func TestFirstFailTooMany(t *testing.T) {

	iterator := &iterstring.Slice{
		Slice: []string{
			"apple",
			"banana",
		},
	}

	var datum string

	if err := (First{iterator}).Decode(&datum); nil == err {
		t.Errorf("Expect an error, but actually did not get one: %v", err)
		return
	} else if expected, actual := errTooManyIterations, err; expected != actual {
		t.Errorf("Expected (%T) %v, but actually got (%T) %v.", expected, expected, actual, actual)
		return
	}
}

func TestFirstFailTooEmpty(t *testing.T) {

	iterator := &iterstring.Slice{
		Slice: []string{
			// Nothing here.
		},
	}

	var datum string

	if err := (First{iterator}).Decode(&datum); nil == err {
		t.Errorf("Expect an error, but actually did not get one: %v", err)
		return
	} else if expected, actual := errEmptyIterator, err; expected != actual {
		t.Errorf("Expected (%T) %v, but actually got (%T) %v.", expected, expected, actual, actual)
		return
	}
}
