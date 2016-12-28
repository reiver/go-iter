package iterutf8

import (
	"bytes"

	"testing"
)

var (
	utf8TestSlices = [][]rune{
		[]rune{},

		[]rune{'0'},
		[]rune{'1'},
		[]rune{'2'},
		[]rune{'3'},
		[]rune{'4'},
		[]rune{'5'},
		[]rune{'6'},
		[]rune{'7'},
		[]rune{'8'},
		[]rune{'9'},

		[]rune{'0','1','2','3','4','5','6','7','8','9'},

		[]rune{'a','p','p','l','e'},
		[]rune{'b','a','n','a','n','a'},
		[]rune{'c','h','e','r','y','y'},

		[]rune{'a','p','p','l','e', ' ', 'b','a','n','a','n','a', ' ', 'c','h','e','r','y','y'},

		[]rune{'O','N','E'},
		[]rune{'T','W','O'},
		[]rune{'T','H','R','E','E'},

		[]rune{'O','N','E', ' ', 'T','W','O', ' ', 'T','H','R','E','E'},

		[]rune{'🙂'},

		[]rune{'💀','👽'},
	}
)

func TestSlice(t *testing.T) {

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

		var actualData []rune
		iterationNumber := -1
		for iterator.Next() {
			iterationNumber++


			var datum rune

			if err := iterator.Decode(&datum); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}

			actualData = append(actualData, datum)


			if err := iterator.Decode((*rune)(nil)); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}


			{
				var x interface{}

				if err := iterator.Decode(&x); nil != err {
					t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
					continue
				}

				datum2, ok := x.(rune)
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

func TestSliceErrNilReceiver(t *testing.T) {

	iterator := (*Slice)(nil)

	{
		err := iterator.Close()
		if nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		}
		if expected, actual := errNilReceiver, err; expected != actual {
			t.Errorf("Expected an (%T) %q, but actually got (%T) %q.", expected, expected, actual, actual)
			return
		}
	}

	{
		var datum interface{}

		err := iterator.Decode(&datum)
		if nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		}
		if expected, actual := errNilReceiver, err; expected != actual {
			t.Errorf("Expected an (%T) %q, but actually got (%T) %q.", expected, expected, actual, actual)
			return
		}
	}

	{
		err := iterator.Err()
		if nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		}
		if expected, actual := errNilReceiver, err; expected != actual {
			t.Errorf("Expected an (%T) %q, but actually got (%T) %q.", expected, expected, actual, actual)
			return
		}
	}

	{
		actual := iterator.Next()
		if expected := false; expected != actual {
			t.Errorf("Expected an %t, but actually got %t.", expected, actual)
			return
		}
	}

	{
		var buffer bytes.Buffer

		_, err := iterator.WriteTo(&buffer)
		if nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		}
		if expected, actual := errNilReceiver, err; expected != actual {
			t.Errorf("Expected an (%T) %q, but actually got (%T) %q.", expected, expected, actual, actual)
			return
		}
	}
}
