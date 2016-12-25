package iter

import (
	"testing"
)

func TestBytes(t *testing.T) {

	tests := []struct{
		Slice []byte
	}{
		{
			Slice: []byte{},
		},



		{
			Slice: []byte{0},
		},
		{
			Slice: []byte{1},
		},
		{
			Slice: []byte{2},
		},
		{
			Slice: []byte{3},
		},
		{
			Slice: []byte{4},
		},
		{
			Slice: []byte{5},
		},
		{
			Slice: []byte{6},
		},
		{
			Slice: []byte{7},
		},
		{
			Slice: []byte{8},
		},
		{
			Slice: []byte{9},
		},
		{
			Slice: []byte{10},
		},



		{
			Slice: []byte{0,1,2,3,4,5,6,7,8,9,10},
		},



		{
			Slice: []byte{213,18,4},
		},
	}


	for testNumber, test := range tests {

		slice := append([]byte(nil), test.Slice...)

		iterator := Bytes{
			Slice: slice,
		}

		if err := iterator.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		var actualData []byte
		iterationNumber := -1
		for iterator.Next() {
			iterationNumber++


			var datum byte

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
