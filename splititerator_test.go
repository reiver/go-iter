package iter_test

import (
	"testing"

	"reflect"
	"strings"

	"github.com/reiver/go-iter"
)

func TestSplitIterator_string(t *testing.T) {

	tests := []struct{
		Strings []string
		Expected [][]string
	}{
		{
			Strings: []string{
				"zero one two",
			},
			Expected: [][]string{
				[]string{
					"zero",
					"one",
					"two",
				},
			},
		},
		{
			Strings: []string{
				"zero one two",
				"THREE FOUR",
			},
			Expected: [][]string{
				[]string{
					"zero",
					"one",
					"two",
				},
				[]string{
					"THREE",
					"FOUR",
				},
			},
		},
		{
			Strings: []string{
				"zero one two",
				"THREE FOUR",
				"five",
			},
			Expected: [][]string{
				[]string{
					"zero",
					"one",
					"two",
				},
				[]string{
					"THREE",
					"FOUR",
				},
				[]string{
					"five",
				},
			},
		},
	}

	fn := func(value string) (iter.Iterator, error) {
		var slice []string = strings.Split(value, " ")
		var iterator iter.Iterator = &iter.Slice[string]{Slice:slice}
		return iterator, nil
	}

	testloop: for testNumber, test := range tests {

		var innerIterator iter.Iterator = &iter.Slice[string]{Slice: test.Strings}

		var iterators iter.Iterators = &iter.SplitIterator[string]{
			Iterator: innerIterator,
			Func: fn,
		}

		var actual [][]string
		iteratorsloop: for {
			var iterator iter.Iterator
			var err error

			iterator, err = iterators.NextIterator()
			if nil != err {
				t.Errorf("For test #%d, did not expect an error but actually got one,", testNumber)
				t.Logf("ERROR: (%T) %s", err, err)
				continue testloop
			}
			if nil == iterator {
				break iteratorsloop
			}

			var slice []string
			for iterator.Next() {
				var dst string

				err := iterator.Decode(&dst)
				if nil != err {
					t.Errorf("For test #%d, did not expect an error but actually got one,", testNumber)
					t.Logf("ERROR: (%T) %s", err, err)
					continue testloop
				}

				slice = append(slice, dst)
			}
			if err := iterator.Err(); nil != err {
				t.Errorf("For test #%d, did not expect an error but actually got one,", testNumber)
				t.Logf("ERROR: (%T) %s", err, err)
				continue testloop
			}

			actual = append(actual, slice)
		}

		{
			expected := test.Expected

			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("For test #%d, the actual value is not what was expected", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				continue testloop
			}
		}
	}
}

