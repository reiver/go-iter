package iter

import (
	"testing"

	"reflect"
)

func TestSlices_nextSlice_string(t *testing.T) {

	tests := []struct{
		Slices Slices[string]
		ExpectedSlice []string
		ExpectedSlices [][]string
	}{
		{
			Slices: Slices[string]{
				Slices: [][]string{},
			},
			ExpectedSlice: nil,
			ExpectedSlices: nil,
		},



		{
			Slices: Slices[string]{
				Slices: [][]string{
					[]string{
						"zero",
					},
				},
			},
			ExpectedSlice: []string{
						"zero",
			},
			ExpectedSlices: nil,
		},
		{
			Slices: Slices[string]{
				Slices: [][]string{
					[]string{
						"zero",
						"one",
					},
				},
			},
			ExpectedSlice: []string{
						"zero",
						"one",
			},
			ExpectedSlices: nil,
		},
		{
			Slices: Slices[string]{
				Slices: [][]string{
					[]string{
						"zero",
						"one",
						"two",
					},
				},
			},
			ExpectedSlice: []string{
						"zero",
						"one",
						"two",
			},
			ExpectedSlices: nil,
		},
		{
			Slices: Slices[string]{
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
			},
			ExpectedSlice: []string{
						"zero",
						"one",
						"two",
			},
			ExpectedSlices: [][]string{
					[]string{
						"THREE",
					},
			},
		},
		{
			Slices: Slices[string]{
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
			},
			ExpectedSlice: []string{
						"zero",
						"one",
						"two",
			},
			ExpectedSlices: [][]string{
					[]string{
						"THREE",
						"FOUR",
					},
			},
		},
		{
			Slices: Slices[string]{
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
			},
			ExpectedSlice: []string{
						"zero",
						"one",
						"two",
			},
			ExpectedSlices: [][]string{
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

	for testNumber, test := range tests {

		if nil == test.Slices.Slices {
			t.Errorf("For test #%d, did not expect inner-slices to be nil but actually was.", testNumber)
			continue
		}

		var actualSlice []string
		{
			var err error

			actualSlice, err = test.Slices.nextSlice()
			if nil != err {
				t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %s", err, err)
				continue
			}
		}

		{
			expected := test.ExpectedSlices
			actual := test.Slices.Slices

			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("For test #%d, the actual slices is not what was expected." , testNumber)
				t.Logf("EXPECTED: (%d) ...", len(expected))
				for i, v := range expected {
					t.Logf("\t[%d] %#v", i, v)
				}
				t.Logf("ACTUAL:   (%d) ...", len(actual))
				for i, v := range actual {
					t.Logf("\t[%d] %#v", i, v)
				}
				continue
			}
		}

		{
			err := test.Slices.Close()
			if nil != err {
				t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %s", err, err)
				continue
			}

			if nil != test.Slices.Slices {
				t.Errorf("For test #%d, expected inner-slices to be nil but actually wasn't.", testNumber)
				continue
			}
		}

		{
			expected := test.ExpectedSlice
			actual   :=        actualSlice

			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("For test #%d, the actual slice is not what was expected." , testNumber)
				t.Logf("EXPECTED: (%d) ...", len(expected))
				for i, v := range expected {
					t.Logf("\t[%d] %q", i, v)
				}
				t.Logf("ACTUAL:   (%d) ...", len(actual))
				for i, v := range actual {
					t.Logf("\t[%d] %q", i, v)
				}
				continue
			}
		}
	}
}
