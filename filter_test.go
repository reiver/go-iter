package iter_test

import (
	"testing"

	"slices"

	"github.com/reiver/go-iter"
)

func TestFilter_string(t *testing.T) {
	tests := []struct{
		Iterator iter.Iterator
		Expected []string
	}{
		{
			Iterator: &iter.Filter[string]{
				Iterator: &iter.Slice[string]{
					Slice:[]string{
				"a",
				"ab",
				"abc",
				"abcd",
				"abcde",
				"abcdef",
				"abcdefg",
				"abcdefgh",
				"abcdefghi",
				"abcdefghij",
				"abcdefghijk",
				"abcdefghijkl",
				"abcdefghijklm",
				"abcdefghijklmn",
				"abcdefghijklmno",
				"abcdefghijklmnop",
				"abcdefghijklmnopq",
				"abcdefghijklmnopqr",
				"abcdefghijklmnopqrs",
				"abcdefghijklmnopqrst",
				"abcdefghijklmnopqrstu",
				"abcdefghijklmnopqrstuv",
				"abcdefghijklmnopqrstuvw",
				"abcdefghijklmnopqrstuvwx",
				"abcdefghijklmnopqrstuvwxy",
				"abcdefghijklmnopqrstuvwxyz",
					},
				},
				Func: func(value string) bool {
					return 5 <= len(value)
				},
			},
			Expected: []string{
				"abcde",
				"abcdef",
				"abcdefg",
				"abcdefgh",
				"abcdefghi",
				"abcdefghij",
				"abcdefghijk",
				"abcdefghijkl",
				"abcdefghijklm",
				"abcdefghijklmn",
				"abcdefghijklmno",
				"abcdefghijklmnop",
				"abcdefghijklmnopq",
				"abcdefghijklmnopqr",
				"abcdefghijklmnopqrs",
				"abcdefghijklmnopqrst",
				"abcdefghijklmnopqrstu",
				"abcdefghijklmnopqrstuv",
				"abcdefghijklmnopqrstuvw",
				"abcdefghijklmnopqrstuvwx",
				"abcdefghijklmnopqrstuvwxy",
				"abcdefghijklmnopqrstuvwxyz",
			},
		},
	}

	testloop: for testNumber, test := range tests {

		var actual []string

		for test.Iterator.Next() {
			var decoded string

			err := test.Iterator.Decode(&decoded)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %s", err, err)
				continue testloop
			}

			actual = append(actual, decoded)
		}
		if err := test.Iterator.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue testloop
		}

		{
			expected := test.Expected

			if !slices.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual slice is not what was expected." , testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				continue
			}
		}

	}
}
