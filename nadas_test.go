package iter

import (
	"testing"
)

func TestNadas(t *testing.T) {

	tests := []struct{
		Slice []struct{}
	}{
		{
			Slice: []struct{}{
				// Nothing here.
			},
		},
		{
			Slice: []struct{}{
				struct{}{},
			},
		},
		{
			Slice: []struct{}{
				struct{}{},
				struct{}{},
			},
		},
		{
			Slice: []struct{}{
				struct{}{},
				struct{}{},
				struct{}{},
			},
		},
		{
			Slice: []struct{}{
				struct{}{},
				struct{}{},
				struct{}{},
				struct{}{},
			},
		},
		{
			Slice: []struct{}{
				struct{}{},
				struct{}{},
				struct{}{},
				struct{}{},
				struct{}{},
			},
		},
	}


	for testNumber, test := range tests {

		slice := append([]struct{}(nil), test.Slice...)

		iterator := Nadas{
			Slice: slice,
		}

		if err := iterator.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		var actualData []struct{}
		iterationNumber := -1
		for iterator.Next() {
			iterationNumber++


			var datum struct{}

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
