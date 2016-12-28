package iterstring

import (
	"fmt"

	"testing"
)

func TestSliceScanIntoString(t *testing.T) {

	tests := []struct{
		Slice []string
	}{}

	for _, slice := range stringTestSlices {
		sliceCopy := append([]string(nil), slice...)

		test := struct{
			Slice []string
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		slice := append([]string(nil), test.Slice...)

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

			var datum string

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
		Slice []string
	}{}

	for _, slice := range stringTestSlices {
		sliceCopy := append([]string(nil), slice...)

		test := struct{
			Slice []string
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		slice := append([]string(nil), test.Slice...)

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

			datum, ok := x.(string)
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
		Slice []string
	}{}

	for _, slice := range stringTestSlices {
		sliceCopy := append([]string(nil), slice...)

		test := struct{
			Slice []string
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		slice := append([]string(nil), test.Slice...)

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

			var datum stringTestScanner

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

type stringTestScanner struct {
	value string
	scanned bool
}

func (receiver stringTestScanner) Scanned() bool {
	return receiver.scanned
}

func (receiver stringTestScanner) Value() string {
	return receiver.value
}

func (receiver *stringTestScanner) Scan(src interface{}) error {
	if nil == src {
		return fmt.Errorf("Cannot scan into nil: (%T) %v", src, src)
	}

	b, ok := src.(string)
	if !ok {
		return fmt.Errorf("Could not convert %T into string.", src)
	}

	receiver.value   = b
	receiver.scanned = true


	return nil
}
