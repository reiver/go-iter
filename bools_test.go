package iter

import (
	"testing"
)

func TestBools(t *testing.T) {

	tests := []struct{
		Slice []bool
	}{
		{
			Slice: []bool{},
		},



		{
			Slice: []bool{false},
		},
		{
			Slice: []bool{true},
		},



		{
			Slice: []bool{false,false},
		},
		{
			Slice: []bool{false,true},
		},
		{
			Slice: []bool{true,false},
		},
		{
			Slice: []bool{true,true},
		},



		{
			Slice: []bool{false,false,false},
		},
		{
			Slice: []bool{false,false,true},
		},
		{
			Slice: []bool{false,true,false},
		},
		{
			Slice: []bool{false,true,true},
		},
		{
			Slice: []bool{true,false,false},
		},
		{
			Slice: []bool{true,false,true},
		},
		{
			Slice: []bool{true,true,false},
		},
		{
			Slice: []bool{true,true,true},
		},



		{
			Slice: []bool{true,false,false,true,false,true,true,false,true,true,true},
		},
	}


	for testNumber, test := range tests {

		slice := append([]bool(nil), test.Slice...)

		iterator := Bools{
			Slice: slice,
		}

		if err := iterator.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		var actualData []bool
		iterationNumber := -1
		for iterator.Next() {
			iterationNumber++


			var datum bool

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
