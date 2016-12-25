package iter

import (
	"testing"
)

func TestComplex64s(t *testing.T) {

	tests := []struct{
		Slice []complex64
	}{
		{
			Slice: []complex64{},
		},



		{
			Slice: []complex64{0},
		},
		{
			Slice: []complex64{1},
		},
		{
			Slice: []complex64{2},
		},
		{
			Slice: []complex64{3},
		},
		{
			Slice: []complex64{4},
		},
		{
			Slice: []complex64{5},
		},
		{
			Slice: []complex64{6},
		},
		{
			Slice: []complex64{7},
		},
		{
			Slice: []complex64{8},
		},
		{
			Slice: []complex64{9},
		},
		{
			Slice: []complex64{10},
		},



		{
			Slice: []complex64{0,1,2,3,4,5,6,7,8,9,10},
		},



		{
			Slice: []complex64{213,18,4},
		},
	}


	for testNumber, test := range tests {

		slice := append([]complex64(nil), test.Slice...)

		iterator := Complex64s{
			Slice: slice,
		}

		if err := iterator.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		var actualData []complex64
		iterationNumber := -1
		for iterator.Next() {
			iterationNumber++


			var datum complex64

			if err := iterator.Decode(&datum); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}

			actualData = append(actualData, datum)
		}
		if err := iterator.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := len(test.Slice), len(actualData); expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}
		for datumNumber, expected := range test.Slice {
			actual := actualData[datumNumber]

			if expected != actual {
				t.Errorf("For test #%d and datum #%d, expected %v, but actually got %v.", testNumber, datumNumber, expected, actual)
				continue
			}
		}

		if err := iterator.Close(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}
	}
}
