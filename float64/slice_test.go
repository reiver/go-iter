package iterfloat64

import (
	"testing"
)

var (
	float64TestSlices = [][]float64{
		[]float64{},

		[]float64{0.0},
		[]float64{1.0},
		[]float64{2.0},
		[]float64{3.0},
		[]float64{4.0},
		[]float64{5.0},
		[]float64{6.0},
		[]float64{7.0},
		[]float64{8.0},
		[]float64{9.0},
		[]float64{10.0},

		[]float64{0.0,1.0,2.0,3.0,4.0,5.0,6.0,7.0,8.0,9.0,10.0},

		[]float64{213.202,18.179,4.000002},
	}
)

func TestSlice(t *testing.T) {

	tests := []struct{
		Slice []float64
	}{}

	for _, slice := range float64TestSlices {
		sliceCopy := append([]float64(nil), slice...)

		test := struct{
			Slice []float64
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		slice := append([]float64(nil), test.Slice...)

		iterator := Slice{
			Slice: slice,
		}

		if err := iterator.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		var actualData []float64
		iterationNumber := -1
		for iterator.Next() {
			iterationNumber++


			var datum float64

			if err := iterator.Decode(&datum); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}

			actualData = append(actualData, datum)


			if err := iterator.Decode((*float64)(nil)); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}


			{
				var x interface{}

				if err := iterator.Decode(&x); nil != err {
					t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
					continue
				}

				datum2, ok := x.(float64)
				if !ok {
					t.Errorf("For test #%d and iteration #%d, expected to be able to cast, but actually could not. (%T)", testNumber, iterationNumber, x)
					continue
				}

				if expected, actual := datum, datum2; expected != actual {
					t.Errorf("For test #%d and iteration #%d, expected %v, but actually got %v.", testNumber, iterationNumber, expected, actual)
					continue
				}
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
