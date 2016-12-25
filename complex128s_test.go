package iter

import (
	"testing"
)

func TestComplex128s(t *testing.T) {

	tests := []struct{
		Slice []complex128
	}{
		{
			Slice: []complex128{},
		},



		{
			Slice: []complex128{0},
		},
		{
			Slice: []complex128{1},
		},
		{
			Slice: []complex128{2},
		},
		{
			Slice: []complex128{3},
		},
		{
			Slice: []complex128{4},
		},
		{
			Slice: []complex128{5},
		},
		{
			Slice: []complex128{6},
		},
		{
			Slice: []complex128{7},
		},
		{
			Slice: []complex128{8},
		},
		{
			Slice: []complex128{9},
		},
		{
			Slice: []complex128{10},
		},



		{
			Slice: []complex128{0,1,2,3,4,5,6,7,8,9,10},
		},



		{
			Slice: []complex128{213,18,4},
		},
	}


	for testNumber, test := range tests {

		slice := append([]complex128(nil), test.Slice...)

		iterator := Complex128s{
			Slice: slice,
		}

		if err := iterator.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		var actualData []complex128
		iterationNumber := -1
		for iterator.Next() {
			iterationNumber++


			var datum complex128

			if err := iterator.Decode(&datum); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}

			actualData = append(actualData, datum)


			if err := iterator.Decode((*complex128)(nil)); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}


			var x interface{}

			if err := iterator.Decode(&x); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}

			datum2, ok := x.(complex128)
			if !ok {
				t.Errorf("For test #%d and iteration #%d, expected to be able to cast, but actually could not. (%T)", testNumber, iterationNumber, x)
				continue
			}

			if expected, actual := datum, datum2; expected != actual {
				t.Errorf("For test #%d and iteration #%d, expected %v, but actually got %v.", testNumber, iterationNumber, expected, actual)
				continue
			}


			if err := iterator.Decode((*interface{})(nil)); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}
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