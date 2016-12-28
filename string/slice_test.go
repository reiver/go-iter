package iterstring

import (
	"bytes"

	"testing"
)

var (
	stringTestSlices = [][]string{
		[]string{},

		[]string{"apple"},
		[]string{"banana"},
		[]string{"cherry"},

		[]string{"ONE"},
		[]string{"TWO"},
		[]string{"THREE"},
		[]string{"FOUR"},
		[]string{"FIVE"},

		[]string{"apple","banana","cherry"},

		[]string{"ONE","TWO","THREE","FOUR","FIVE"},
	}
)

func TestSlice(t *testing.T) {

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

		var actualData []string
		iterationNumber := -1
		for iterator.Next() {
			iterationNumber++


			var datum string

			if err := iterator.Decode(&datum); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}

			actualData = append(actualData, datum)


			if err := iterator.Decode((*string)(nil)); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}


			{
				var x interface{}

				if err := iterator.Decode(&x); nil != err {
					t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
					continue
				}

				datum2, ok := x.(string)
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

func TestSliceClose(t *testing.T) {

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

		for closeTestNumber:=0; closeTestNumber<len(test.Slice); closeTestNumber++ {
			slice := append([]string(nil), test.Slice...)

			iterator := Slice{
				Slice: slice,
			}

			for i:=0; i<closeTestNumber; i++ {
				if expected, actual := true, iterator.Next(); expected != actual {
					t.Errorf("For test #%d and close test #%d, expected %t, but actually got %t.", testNumber, closeTestNumber, expected, actual)
					continue
				}
			}

			if err := iterator.Close(); nil != err {
				t.Errorf("For test #%d, and close test #%d did not expect an error, but actually got one: (%T) %v", testNumber, closeTestNumber, err, err)
				continue
			}

			if expected, actual := false, iterator.Next(); expected != actual {
				t.Errorf("For test #%d and close test #%d, expected %t, but actually got %t.", testNumber, closeTestNumber, expected, actual)
				continue
			}

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
