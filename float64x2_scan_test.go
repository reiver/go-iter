package iter

import (
	"testing"
)

func TestFloat64x2ScanIntoFloat64x2(t *testing.T) {

	tests := []struct{
		Slice [][2]float64
	}{}

	for _, slice := range float64x2TestSlices {
		sliceCopy := append([][2]float64(nil), slice...)

		test := struct{
			Slice [][2]float64
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		slice := append([][2]float64(nil), test.Slice...)

		iterator := Float64x2{
			Slice: slice,
		}

		if err := iterator.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		iterationNumber := -1
		for iterator.Next() {
			iterationNumber++

			var datum [2]float64

			if err := iterator.Scan(&datum[0], &datum[1]); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}

			if expected, actual := test.Slice[iterationNumber], datum; expected != actual {
				t.Errorf("For test #%d and iteration #%d, expected %v, but actually got %v.", testNumber, iterationNumber, expected, actual)
				continue
			}
		}
	}
}

func TestFloat64x2ScanIntoInterface(t *testing.T) {

	tests := []struct{
		Slice [][2]float64
	}{}

	for _, slice := range float64x2TestSlices {
		sliceCopy := append([][2]float64(nil), slice...)

		test := struct{
			Slice [][2]float64
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		slice := append([][2]float64(nil), test.Slice...)

		iterator := Float64x2{
			Slice: slice,
		}

		if err := iterator.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		iterationNumber := -1
		for iterator.Next() {
			iterationNumber++

			var x0 interface{}
			var x1 interface{}

			if err := iterator.Scan(&x0, &x1); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}

			datum0, ok := x0.(float64)
			if !ok {
				t.Errorf("For test #%d and iteration #%d, expected to be able to cast, but actually couldn't.", testNumber, iterationNumber)
				continue
			}
			datum1, ok := x1.(float64)
			if !ok {
				t.Errorf("For test #%d and iteration #%d, expected to be able to cast, but actually couldn't.", testNumber, iterationNumber)
				continue
			}

			if expected, actual := test.Slice[iterationNumber][0], datum0; expected != actual {
				t.Errorf("For test #%d and iteration #%d, expected %v, but actually got %v.", testNumber, iterationNumber, expected, actual)
				continue
			}
			if expected, actual := test.Slice[iterationNumber][1], datum1; expected != actual {
				t.Errorf("For test #%d and iteration #%d, expected %v, but actually got %v.", testNumber, iterationNumber, expected, actual)
				continue
			}
		}
	}
}

func TestFloat64x2ScanIntoScanner(t *testing.T) {

	tests := []struct{
		Slice [][2]float64
	}{}

	for _, slice := range float64x2TestSlices {
		sliceCopy := append([][2]float64(nil), slice...)

		test := struct{
			Slice [][2]float64
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		slice := append([][2]float64(nil), test.Slice...)

		iterator := Float64x2{
			Slice: slice,
		}

		if err := iterator.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		iterationNumber := -1
		for iterator.Next() {
			iterationNumber++

			var datum0 float64TestScanner
			var datum1 float64TestScanner

			if err := iterator.Scan(&datum0, &datum1); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}

			if !datum0.Scanned() {
				t.Errorf("For test #%d and iteration #%d, expected Scan to work, but actually didn't.", testNumber, iterationNumber)
				continue
			}
			if !datum1.Scanned() {
				t.Errorf("For test #%d and iteration #%d, expected Scan to work, but actually didn't.", testNumber, iterationNumber)
				continue
			}

			if expected, actual := test.Slice[iterationNumber][0], datum0.Value(); expected != actual {
				t.Errorf("For test #%d and iteration #%d, expected %v, but actually got %v.", testNumber, iterationNumber, expected, actual)
				continue
			}
			if expected, actual := test.Slice[iterationNumber][1], datum1.Value(); expected != actual {
				t.Errorf("For test #%d and iteration #%d, expected %v, but actually got %v.", testNumber, iterationNumber, expected, actual)
				continue
			}
		}
	}
}
