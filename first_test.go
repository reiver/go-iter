package iter

import (
	"github.com/reiver/go-iter/string"

	"testing"
)

func TestFirst(t *testing.T) {

	tests := []struct{
		Iterator Iterator
		Expected string
	}{
		{
			Iterator: &iterstring.Slice{
				Slice:[]string{
					"apple",
				},
			},
			Expected: "apple",
		},
		{
			Iterator: &iterstring.Slice{
				Slice:[]string{
					"apple",
					"banana",
				},
			},
			Expected: "apple",
		},
		{
			Iterator: &iterstring.Slice{
				Slice:[]string{
					"apple",
					"banana",
					"cherry",
				},
			},
			Expected: "apple",
		},
	}

	for testNumber, test := range tests {

		var actual string

		if err := (First{test.Iterator}).Decode(&actual); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		{
			expected := test.Expected
			if expected != actual {
				t.Errorf("For test #%d, the actual first value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				continue
			}
		}


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
