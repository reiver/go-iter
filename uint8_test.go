package iter

import (
	"testing"
)

func TestUint8(t *testing.T) {

	tests := []struct{
		Slice []uint8
	}{
		{
			Slice: []uint8{},
		},



		{
			Slice: []uint8{0},
		},
		{
			Slice: []uint8{1},
		},
		{
			Slice: []uint8{2},
		},
		{
			Slice: []uint8{3},
		},
		{
			Slice: []uint8{4},
		},
		{
			Slice: []uint8{5},
		},
		{
			Slice: []uint8{6},
		},
		{
			Slice: []uint8{7},
		},
		{
			Slice: []uint8{8},
		},
		{
			Slice: []uint8{9},
		},
		{
			Slice: []uint8{10},
		},



		{
			Slice: []uint8{0,1,2,3,4,5,6,7,8,9,10},
		},



		{
			Slice: []uint8{213,18,4},
		},
	}


	for testNumber, test := range tests {

		slice := append([]uint8(nil), test.Slice...)

		iterator := Uint8{
			Slice: slice,
		}

		if err := iterator.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		var actualData []uint8
		iterationNumber := -1
		for iterator.Next() {
			iterationNumber++


			var datum uint8

			if err := iterator.Decode(&datum); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}

			actualData = append(actualData, datum)


			if err := iterator.Decode((*uint8)(nil)); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}


			var x interface{}

			if err := iterator.Decode(&x); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}

			datum2, ok := x.(uint8)
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
