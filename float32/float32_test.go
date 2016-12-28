package iterfloat32

import (
	"testing"
)

var (
	float32TestSlices = [][]float32{
		[]float32{},

		[]float32{0.0},
		[]float32{1.0},
		[]float32{2.0},
		[]float32{3.0},
		[]float32{4.0},
		[]float32{5.0},
		[]float32{6.0},
		[]float32{7.0},
		[]float32{8.0},
		[]float32{9.0},
		[]float32{10.0},

		[]float32{0.0,1.0,2.0,3.0,4.0,5.0,6.0,7.0,8.0,9.0,10.0},

		[]float32{213.202,18.179,4.000002},
	}
)

func TestSlice(t *testing.T) {

	tests := []struct{
		Slice []float32
	}{}

	for _, slice := range float32TestSlices {
		sliceCopy := append([]float32(nil), slice...)

		test := struct{
			Slice []float32
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		slice := append([]float32(nil), test.Slice...)

		iterator := Slice{
			Slice: slice,
		}

		if err := iterator.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		var actualData []float32
		iterationNumber := -1
		for iterator.Next() {
			iterationNumber++


			var datum float32

			if err := iterator.Decode(&datum); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}

			actualData = append(actualData, datum)


			if err := iterator.Decode((*float32)(nil)); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}


			{
				var x interface{}

				if err := iterator.Decode(&x); nil != err {
					t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
					continue
				}

				datum2, ok := x.(float32)
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
		Slice []float32
	}{}

	for _, slice := range float32TestSlices {
		sliceCopy := append([]float32(nil), slice...)

		test := struct{
			Slice []float32
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		for closeTestNumber:=0; closeTestNumber<len(test.Slice); closeTestNumber++ {
			slice := append([]float32(nil), test.Slice...)

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
}
