package iterutf8

import (
	"bytes"
	"unicode/utf8"

	"testing"
)

func TestSliceWriteTo(t *testing.T) {

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

		var buffer bytes.Buffer

		iterationNumber := -1
		for iterator.Next() {
			iterationNumber++

			iterator.WriteTo(&buffer)

			var buffer2 bytes.Buffer
			iterator.WriteTo(&buffer2)
			if expected, actual := utf8.RuneLen(test.Slice[iterationNumber]), buffer2.Len(); expected != actual {
				t.Errorf("For test #%d and iteration number #%d, expected %d, but actually got %d.", testNumber, iterationNumber, expected, actual)
				continue
			}
			{
				actualRune, _ := utf8.DecodeRune(buffer2.Bytes())
				if expected, actual := test.Slice[iterationNumber], actualRune; expected != actual {
					t.Errorf("For test #%d and iteration number #%d, expected %d (%q), but actually got %d (%q).", testNumber, iterationNumber, expected, expected, actual, actual)
					continue
				}
			}
		}
		if err := iterator.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v" , testNumber, err, err)
			continue
		}

		if expected, actual := len(test.Slice), utf8.RuneCount(buffer.Bytes()); expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d." , testNumber, expected, actual)
			continue
		}
		if expected, actual := string(test.Slice), buffer.String(); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q." , testNumber, expected, actual)
			continue
		}
	}
}
