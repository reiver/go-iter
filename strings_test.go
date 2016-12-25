package iter

import (
	"testing"
)

func TestStrings(t *testing.T) {

	tests := []struct{
		Slice []string
	}{
		{
			Slice: []string{},
		},



		{
			Slice: []string{"apple"},
		},
		{
			Slice: []string{"banana"},
		},
		{
			Slice: []string{"cherry"},
		},



		{
			Slice: []string{"ONE"},
		},
		{
			Slice: []string{"TWO"},
		},
		{
			Slice: []string{"THREE"},
		},
		{
			Slice: []string{"FOUR"},
		},
		{
			Slice: []string{"FIVE"},
		},



		{
			Slice: []string{"apple","banana","cherry"},
		},



		{
			Slice: []string{"ONE","TWO","THREE","FOUR","FIVE"},
		},
	}


	for testNumber, test := range tests {

		slice := append([]string(nil), test.Slice...)

		iterator := Strings{
			Slice: slice,
		}

		if err := iterator.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		var actualData []string
		iterationNumber := -1
		for iterator.Next() {
			iterationNumber++


			var datum string

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
