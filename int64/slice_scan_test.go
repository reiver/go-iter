package iterint64

import (
	"fmt"

	"testing"
)

func TestSliceScanIntoInt64(t *testing.T) {

	tests := []struct{
		Slice []int64
	}{}

	for _, slice := range int64TestSlices {
		sliceCopy := append([]int64(nil), slice...)

		test := struct{
			Slice []int64
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		slice := append([]int64(nil), test.Slice...)

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

			var datum int64

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

func TestSliceScanIntoInterface(t *testing.T) {

	tests := []struct{
		Slice []int64
	}{}

	for _, slice := range int64TestSlices {
		sliceCopy := append([]int64(nil), slice...)

		test := struct{
			Slice []int64
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		slice := append([]int64(nil), test.Slice...)

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

			datum, ok := x.(int64)
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

func TestSliceScanIntoScanner(t *testing.T) {

	tests := []struct{
		Slice []int64
	}{}

	for _, slice := range int64TestSlices {
		sliceCopy := append([]int64(nil), slice...)

		test := struct{
			Slice []int64
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		slice := append([]int64(nil), test.Slice...)

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

			var datum int64TestScanner

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

type int64TestScanner struct {
	value int64
	scanned bool
}

func (receiver int64TestScanner) Scanned() bool {
	return receiver.scanned
}

func (receiver int64TestScanner) Value() int64 {
	return receiver.value
}

func (receiver *int64TestScanner) Scan(src interface{}) error {
	if nil == src {
		return fmt.Errorf("Cannot scan into nil: (%T) %v", src, src)
	}

	b, ok := src.(int64)
	if !ok {
		return fmt.Errorf("Could not convert %T into int64.", src)
	}

	receiver.value   = b
	receiver.scanned = true


	return nil
}
