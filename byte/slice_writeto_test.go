package iterbyte

import (
	"bytes"

	"testing"
)

func TestSliceWriteTo(t *testing.T) {

	tests := []struct{
		Slice []byte
	}{
		{
			Slice: []byte(nil),
		},
		{
			Slice: []byte{},
		},



		{
			Slice: []byte("apple"),
		},
		{
			Slice: []byte("banana"),
		},
		{
			Slice: []byte("cherry"),
		},



		{
			Slice: []byte("apple banana cherry"),
		},



		{
			Slice: []byte("ONE"),
		},
		{
			Slice: []byte("TWO"),
		},
		{
			Slice: []byte("THREE"),
		},



		{
			Slice: []byte("ONE TWO THREE"),
		},



		{
			Slice: []byte("😐"),
		},



		{
			Slice: []byte("💀👽"),
		},

	}


	for testNumber, test := range tests {

		slice := append([]byte(nil), test.Slice...)

		iterator := Slice{
			Slice: slice,
		}

		var buffer bytes.Buffer

		iterationNumber := -1
		for iterator.Next() {
			iterationNumber++

			iterator.WriteTo(&buffer)

			var buffer2 bytes.Buffer
			iterator.WriteTo(&buffer2)
			if expected, actual := 1, buffer2.Len(); expected != actual {
				t.Errorf("For test #%d and iteration number #%d, expected %d, but actually got %d.", testNumber, iterationNumber, expected, actual)
				continue
			}
			if expected, actual := test.Slice[iterationNumber], buffer2.Bytes()[0]; expected != actual {
				t.Errorf("For test #%d and iteration number #%d, expected %d (%q), but actually got %d (%q).", testNumber, iterationNumber, expected, expected, actual, actual)
				continue
			}
		}
		if err := iterator.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v" , testNumber, err, err)
			continue
		}

		if expected, actual := len(test.Slice), len(buffer.Bytes()); expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d." , testNumber, expected, actual)
			continue
		}
		if expected, actual := string(test.Slice), buffer.String(); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q." , testNumber, expected, actual)
			continue
		}
	}
}
