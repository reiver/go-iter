package iteruint16

import (
	"fmt"

	"testing"
)

func TestSliceScanIntoInt32(t *testing.T) {

	tests := []struct{
		Slice []uint16
	}{}

	for _, slice := range uint16TestSlices {
		sliceCopy := append([]uint16(nil), slice...)

		test := struct{
			Slice []uint16
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		slice := append([]uint16(nil), test.Slice...)

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

			var datum int32

			if err := iterator.Scan(&datum); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}

			if expected, actual := int32(test.Slice[iterationNumber]), datum; expected != actual {
				t.Errorf("For test #%d and iteration #%d, expected %v, but actually got %v.", testNumber, iterationNumber, expected, actual)
				continue
			}
		}
	}
}

func TestSliceScanIntoInt64(t *testing.T) {

	tests := []struct{
		Slice []uint16
	}{}

	for _, slice := range uint16TestSlices {
		sliceCopy := append([]uint16(nil), slice...)

		test := struct{
			Slice []uint16
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		slice := append([]uint16(nil), test.Slice...)

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

			if expected, actual := int64(test.Slice[iterationNumber]), datum; expected != actual {
				t.Errorf("For test #%d and iteration #%d, expected %v, but actually got %v.", testNumber, iterationNumber, expected, actual)
				continue
			}
		}
	}
}

func TestSliceScanIntoUint16(t *testing.T) {

	tests := []struct{
		Slice []uint16
	}{}

	for _, slice := range uint16TestSlices {
		sliceCopy := append([]uint16(nil), slice...)

		test := struct{
			Slice []uint16
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		slice := append([]uint16(nil), test.Slice...)

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

			var datum uint16

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

func TestSliceScanIntoUint32(t *testing.T) {

	tests := []struct{
		Slice []uint16
	}{}

	for _, slice := range uint16TestSlices {
		sliceCopy := append([]uint16(nil), slice...)

		test := struct{
			Slice []uint16
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		slice := append([]uint16(nil), test.Slice...)

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

			var datum uint32

			if err := iterator.Scan(&datum); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}

			if expected, actual := uint32(test.Slice[iterationNumber]), datum; expected != actual {
				t.Errorf("For test #%d and iteration #%d, expected %v, but actually got %v.", testNumber, iterationNumber, expected, actual)
				continue
			}
		}
	}
}

func TestSliceScanIntoUint64(t *testing.T) {

	tests := []struct{
		Slice []uint16
	}{}

	for _, slice := range uint16TestSlices {
		sliceCopy := append([]uint16(nil), slice...)

		test := struct{
			Slice []uint16
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		slice := append([]uint16(nil), test.Slice...)

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

			var datum uint64

			if err := iterator.Scan(&datum); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}

			if expected, actual := uint64(test.Slice[iterationNumber]), datum; expected != actual {
				t.Errorf("For test #%d and iteration #%d, expected %v, but actually got %v.", testNumber, iterationNumber, expected, actual)
				continue
			}
		}
	}
}

func TestSliceScanIntoInterface(t *testing.T) {

	tests := []struct{
		Slice []uint16
	}{}

	for _, slice := range uint16TestSlices {
		sliceCopy := append([]uint16(nil), slice...)

		test := struct{
			Slice []uint16
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		slice := append([]uint16(nil), test.Slice...)

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

			datum, ok := x.(uint16)
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
		Slice []uint16
	}{}

	for _, slice := range uint16TestSlices {
		sliceCopy := append([]uint16(nil), slice...)

		test := struct{
			Slice []uint16
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		slice := append([]uint16(nil), test.Slice...)

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

			var datum uint16TestScanner

			if err := iterator.Scan(&datum); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}

			if !datum.Scanned() {
				t.Errorf("For test #%d and iteration #%d, expected Scan to work, but actually didn't.", testNumber, iterationNumber)
				continue
			}

			if expected, actual := int64(test.Slice[iterationNumber]), datum.Value(); expected != actual {
				t.Errorf("For test #%d and iteration #%d, expected %v, but actually got %v.", testNumber, iterationNumber, expected, actual)
				continue
			}
		}
	}
}

type uint16TestScanner struct {
	value int64
	scanned bool
}

func (receiver uint16TestScanner) Scanned() bool {
	return receiver.scanned
}

func (receiver uint16TestScanner) Value() int64 {
	return receiver.value
}

func (receiver *uint16TestScanner) Scan(src interface{}) error {
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
