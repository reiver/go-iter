package iterutf8

import (
	"bytes"
	"fmt"

	"testing"
)

func TestSliceScanIntoRune(t *testing.T) {

	tests := []struct{
		Slice []rune
	}{}

	for _, slice := range utf8TestSlices {
		sliceCopy := append([]rune(nil), slice...)

		test := struct{
			Slice []rune
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		var slice []byte
		{
			var temp bytes.Buffer

			for _, r := range test.Slice {
				temp.WriteRune(r)
			}

			slice = append([]byte(nil), temp.Bytes()...)
		}

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

			var datum rune

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
		Slice []rune
	}{}

	for _, slice := range utf8TestSlices {
		sliceCopy := append([]rune(nil), slice...)

		test := struct{
			Slice []rune
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		var slice []byte
		{
			var temp bytes.Buffer

			for _, r := range test.Slice {
				temp.WriteRune(r)
			}

			slice = append([]byte(nil), temp.Bytes()...)
		}

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

			datum, ok := x.(rune)
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
		Slice []rune
	}{}

	for _, slice := range utf8TestSlices {
		sliceCopy := append([]rune(nil), slice...)

		test := struct{
			Slice []rune
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		var slice []byte
		{
			var temp bytes.Buffer

			for _, r := range test.Slice {
				temp.WriteRune(r)
			}

			slice = append([]byte(nil), temp.Bytes()...)
		}

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

			var datum utf8TestScanner

			if err := iterator.Scan(&datum); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}

			if !datum.Scanned() {
				t.Errorf("For test #%d and iteration #%d, expected Scan to work, but actually didn't.", testNumber, iterationNumber)
				continue
			}

			if expected, actual := string(test.Slice[iterationNumber]), string(datum.Value()); expected != actual {
				t.Errorf("For test #%d and iteration #%d, expected %v, but actually got %v.", testNumber, iterationNumber, expected, actual)
				continue
			}
		}
	}
}

type utf8TestScanner struct {
	value []byte
	scanned bool
}

func (receiver utf8TestScanner) Scanned() bool {
	return receiver.scanned
}

func (receiver utf8TestScanner) Value() []byte {
	return receiver.value
}

func (receiver *utf8TestScanner) Scan(src interface{}) error {
	if nil == src {
		return fmt.Errorf("Cannot scan into nil: (%T) %v", src, src)
	}

	b, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("Could not convert %T into []byte.", src)
	}

	receiver.value   = b
	receiver.scanned = true


	return nil
}
