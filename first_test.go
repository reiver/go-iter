package iter_test

import (
	"testing"

	"github.com/reiver/go-iter"
)

func TestFirst(t *testing.T) {

	tests := []struct{
		Iterator iter.Iterator
		Expected string
	}{
		{
			Iterator: &iter.Slice[string]{
				Slice:[]string{
					"apple",
				},
			},
			Expected: "apple",
		},
		{
			Iterator: &iter.Slice[string]{
				Slice:[]string{
					"apple",
					"banana",
				},
			},
			Expected: "apple",
		},
		{
			Iterator: &iter.Slice[string]{
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

		if err := (iter.First{test.Iterator}).Decode(&actual); nil != err {
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

	iterator := &iter.Slice[string]{
		Slice: []string{
			// Nothing here.
		},
	}

	var datum string

	if err := (iter.First{iterator}).Decode(&datum); nil == err {
		t.Errorf("Expect an error, but actually did not get one: %v", err)
		return
	} else if expected, actual := "iter: empty iterator", err.Error(); expected != actual {
		t.Errorf("Expected error %q, but actually got error %q.", expected, actual)
		return
	}
}
