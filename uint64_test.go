package iter

import (
	"testing"
)

func TestUint64(t *testing.T) {

	tests := []struct{
		Slice []uint64
	}{
		{
			Slice: []uint64{},
		},



		{
			Slice: []uint64{0},
		},
		{
			Slice: []uint64{1},
		},
		{
			Slice: []uint64{2},
		},
		{
			Slice: []uint64{3},
		},
		{
			Slice: []uint64{4},
		},
		{
			Slice: []uint64{5},
		},
		{
			Slice: []uint64{6},
		},
		{
			Slice: []uint64{7},
		},
		{
			Slice: []uint64{8},
		},
		{
			Slice: []uint64{9},
		},
		{
			Slice: []uint64{10},
		},



		{
			Slice: []uint64{0,1,2,3,4,5,6,7,8,9,10},
		},



		{
			Slice: []uint64{213,18,4},
		},
	}


	for testNumber, test := range tests {

		slice := append([]uint64(nil), test.Slice...)

		iterator := Uint64{
			Slice: slice,
		}

		if err := iterator.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		var actualData []uint64
		iterationNumber := -1
		for iterator.Next() {
			iterationNumber++


			var datum uint64

			if err := iterator.Decode(&datum); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}

			actualData = append(actualData, datum)


			if err := iterator.Decode((*uint64)(nil)); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}


			var x interface{}

			if err := iterator.Decode(&x); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}

			datum2, ok := x.(uint64)
			if !ok {
				t.Errorf("For test #%d and iteration #%d, expected to be able to cast, but actually could not. (%T)", testNumber, iterationNumber, x)
				continue
			}

			if expected, actual := datum, datum2; expected != actual {
				t.Errorf("For test #%d and iteration #%d, expected %v, but actually got %v.", testNumber, iterationNumber, expected, actual)
				continue
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

func TestUint64ErrNilReceiver(t *testing.T) {

	iterator := (*Uint64)(nil)

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
}
