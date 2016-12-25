package iter

import (
	"testing"
)

func TestFloat32s(t *testing.T) {

	tests := []struct{
		Slice []float32
	}{
		{
			Slice: []float32{},
		},



		{
			Slice: []float32{0.0},
		},
		{
			Slice: []float32{1.0},
		},
		{
			Slice: []float32{2.0},
		},
		{
			Slice: []float32{3.0},
		},
		{
			Slice: []float32{4.0},
		},
		{
			Slice: []float32{5.0},
		},
		{
			Slice: []float32{6.0},
		},
		{
			Slice: []float32{7.0},
		},
		{
			Slice: []float32{8.0},
		},
		{
			Slice: []float32{9.0},
		},
		{
			Slice: []float32{10.0},
		},



		{
			Slice: []float32{0.0,1.0,2.0,3.0,4.0,5.0,6.0,7.0,8.0,9.0,10.0},
		},



		{
			Slice: []float32{213.202,18.179,4.000002},
		},
	}


	for testNumber, test := range tests {

		slice := append([]float32(nil), test.Slice...)

		iterator := Float32s{
			Slice: slice,
		}

		if err := iterator.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		var actualData []float32
		iterationNumber := -1
		for iterator.Next() {
			iterationNumber++


			var datum float32

			if err := iterator.Decode(&datum); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}

			actualData = append(actualData, datum)


			var x interface{}

			if err := iterator.Decode(&x); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}

			datum2, ok := x.(float32)
			if !ok {
				t.Errorf("For test #%d and iteration #%d, expected to be able to cast, but actually could not. (%T)", testNumber, iterationNumber, x)
				continue
			}

			if expected, actual := datum, datum2; expected != actual {
				t.Errorf("For test #%d and iteration #%d, expected %v, but actually got %v.", testNumber, iterationNumber, expected, actual)
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
