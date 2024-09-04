package iter_test

import (
	"testing"

	"github.com/reiver/go-iter"
)

func TestSlices_Close_string(t *testing.T) {

	tests := []struct{
		Slices iter.Slices[string]
	}{
		{
			Slices: iter.Slices[string]{
				Slices: [][]string{},
			},
		},



		{
			Slices: iter.Slices[string]{
				Slices: [][]string{
					[]string{
						"zero",
					},
				},
			},
		},
		{
			Slices: iter.Slices[string]{
				Slices: [][]string{
					[]string{
						"zero",
						"one",
					},
				},
			},
		},
		{
			Slices: iter.Slices[string]{
				Slices: [][]string{
					[]string{
						"zero",
						"one",
						"two",
					},
				},
			},
		},
		{
			Slices: iter.Slices[string]{
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
		},
		{
			Slices: iter.Slices[string]{
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
		},
		{
			Slices: iter.Slices[string]{
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
		},
	}

	for testNumber, test := range tests {

		if nil == test.Slices.Slices {
			t.Errorf("For test #%d, did not expect inner-slices to be nil but actually was.", testNumber)
			continue
		}

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
}
