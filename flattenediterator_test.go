package iter_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-iter"
)

func TestFlattenedIterator_string(t *testing.T) {

	tests := []struct{
		Slices [][]string
		Expected []string
	}{
		{
			Slices: [][]string{
				[]string{},
			},
			Expected: nil,
		},
		{
			Slices: [][]string{
				[]string{
					"zero",
				},
			},
			Expected: []string{
				"zero",
			},
		},
		{
			Slices: [][]string{
				[]string{
					"zero",
					"one",
				},
			},
			Expected: []string{
				"zero",
				"one",
			},
		},
		{
			Slices: [][]string{
				[]string{
					"zero",
					"one",
					"two",
				},
			},
			Expected: []string{
				"zero",
				"one",
				"two",
			},
		},
		{
			Slices: [][]string{
				[]string{
					"zero",
					"one",
					"two",
				},
				[]string{
					"THREE",
				},
			},
			Expected: []string{
				"zero",
				"one",
				"two",
				"THREE",
			},
		},
		{
			Slices: [][]string{
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
			Expected: []string{
				"zero",
				"one",
				"two",
				"THREE",
				"FOUR",
			},
		},
		{
			Slices: [][]string{
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
			Expected: []string{
				"zero",
				"one",
				"two",
				"THREE",
				"FOUR",
				"five",
			},
		},
	}

	for testNumber, test := range tests {

		var actual []string

		var iterator iter.FlattenedIterator
		iterator.Iterators = &iter.Slices[string]{
			Slices: test.Slices,
		}

		for iterator.Next() {
			var dst string

			err := iterator.Decode(&dst)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %s", err, err)
				t.Logf("SLICES: %#v", test.Slices)
				continue
			}

			actual = append(actual, dst)
		}
		if err := iterator.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("SLICES: %#v", test.Slices)
			continue
		}

		{
			expected := test.Expected

			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("For test #%d, the actual flatted result is not what was expected.", testNumber)
				t.Logf("SLICES: %#v", test.Slices)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL  : %#v", actual)
				t.Logf("EXPECTED: (%d) ...", len(expected))
				for i, v := range expected {
					t.Logf("\t[%d] %q", i, v)
				}
				t.Logf("ACTUAL:   (%d)...", len(actual))
				for i, v := range actual {
					t.Logf("\t[%d] %q", i, v)
				}
				continue
			}
		}

		{
			err := iterator.Close()
			if nil != err {
				t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %s", err, err)
				t.Logf("SLICES: %#v", test.Slices)
				continue
			}
		}
	}
}
