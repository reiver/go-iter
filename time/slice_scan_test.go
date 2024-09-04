package itertime

import (
	"testing"

	"time"

	"github.com/reiver/go-iter/internal/testing"
)

func TestTimeScanIntoTime(t *testing.T) {

	tests := []struct{
		Slice []time.Time
	}{}

	for _, slice := range timeTestSlices {
		sliceCopy := append([]time.Time(nil), slice...)

		test := struct{
			Slice []time.Time
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		slice := append([]time.Time(nil), test.Slice...)

		iterator := Slice{
			Slice: slice,
		}

		if err := iterator.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		iterationNumber := -1
		for iterator.Next() {
			iterationNumber++

			var datum time.Time

			if err := iterator.Scan(&datum); nil != err {
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

func TestTimeScanIntoInterface(t *testing.T) {

	tests := []struct{
		Slice []time.Time
	}{}

	for _, slice := range timeTestSlices {
		sliceCopy := append([]time.Time(nil), slice...)

		test := struct{
			Slice []time.Time
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		slice := append([]time.Time(nil), test.Slice...)

		iterator := Slice{
			Slice: slice,
		}

		if err := iterator.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		iterationNumber := -1
		for iterator.Next() {
			iterationNumber++

			var x interface{}

			if err := iterator.Scan(&x); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}

			datum, ok := x.(time.Time)
			if !ok {
				t.Errorf("For test #%d and iteration #%d, expected to be able to cast, but actually couldn't.", testNumber, iterationNumber)
				continue
			}

			if expected, actual := test.Slice[iterationNumber], datum; expected != actual {
				t.Errorf("For test #%d and iteration #%d, expected %v, but actually got %v.", testNumber, iterationNumber, expected, actual)
				continue
			}
		}
	}
}

func TestTimeScanIntoScanner(t *testing.T) {

	tests := []struct{
		Slice []time.Time
	}{}

	for _, slice := range timeTestSlices {
		sliceCopy := append([]time.Time(nil), slice...)

		test := struct{
			Slice []time.Time
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		slice := append([]time.Time(nil), test.Slice...)

		iterator := Slice{
			Slice: slice,
		}

		if err := iterator.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		iterationNumber := -1
		for iterator.Next() {
			iterationNumber++

			var datum internaltesting.TestScanner[time.Time]

			if err := iterator.Scan(&datum); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}

			if !datum.Scanned() {
				t.Errorf("For test #%d and iteration #%d, expected Scan to work, but actually didn't.", testNumber, iterationNumber)
				continue
			}

			if expected, actual := test.Slice[iterationNumber], datum.Value(); expected != actual {
				t.Errorf("For test #%d and iteration #%d, expected %v, but actually got %v.", testNumber, iterationNumber, expected, actual)
				continue
			}
		}
	}
}
